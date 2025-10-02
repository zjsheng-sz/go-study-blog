package middleware

import (
	"errors"
	"go-study-blog/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ErroHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err

			switch e := err.(type) {
			case *common.AppError:
				// 使用自定义的错误状态码
				statusCode := http.StatusOK
				if e.Code >= 400 {
					statusCode = e.Code
				}

				ctx.JSON(statusCode, gin.H{
					"success": false,
					"code":    e.Code,
					"message": e.Message,
					"data":    nil,
				})

			case error:
				var statusCode int
				var appErr *common.AppError

				// 处理常见错误类型
				switch {
				case e == gorm.ErrRecordNotFound:
					appErr = common.ErrNotFound
					statusCode = http.StatusNotFound

				case errors.Is(e, bcrypt.ErrMismatchedHashAndPassword):
					appErr = common.ErrInvalidPassword
					statusCode = http.StatusBadRequest

				default:
					appErr = common.ErrInternal
					statusCode = http.StatusInternalServerError
					// 在开发环境下记录详细错误
					if gin.Mode() != gin.ReleaseMode {
						appErr.Message = e.Error()
					}
				}

				ctx.JSON(statusCode, gin.H{
					"success": false,
					"code":    appErr.Code,
					"message": appErr.Message,
					"data":    nil,
				})
			}
		}
	}
}
