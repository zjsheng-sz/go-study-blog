package middleware

import (
	"go-study-blog/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ErroHandler() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err

			switch e := err.(type) {

			case *common.AppError:
				ctx.JSON(http.StatusOK, gin.H{
					"code":    e.Code,
					"message": e.Message,
				})
			case error:
				if e == gorm.ErrRecordNotFound {
					ctx.JSON(http.StatusOK, gin.H{
						"code":    common.ErrNotFound.Code,
						"message": common.ErrNotFound.Message,
					})

				} else {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"code":    common.ErrInternal.Code,
						"message": common.ErrInternal.Message,
					})
				}
			}

		}

	}

}
