package middleware

import (
	"go-study-blog/config"
	"go-study-blog/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		cfg := config.Load()

		authHead := ctx.GetHeader("Authorization")

		if authHead == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHead, "Bearer ")

		if tokenString == authHead {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			ctx.Abort()
			return
		}

		claims, err := utils.ParseToken(tokenString, cfg.App.JWTSecret)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}


