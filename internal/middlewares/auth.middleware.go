package middlewares

import (
	"github.com/chrugez/go-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(ctx, response.ErrCodeInvalidToken, "")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}