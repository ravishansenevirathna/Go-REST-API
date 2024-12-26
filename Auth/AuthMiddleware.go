package Auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/utils"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Extract the token from the Authorization header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			ctx.Abort()
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Validate the token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		// Attach claims to the context for access in handlers
		ctx.Set("claims", claims)

		// Proceed to the next handler
		ctx.Next()
	}
}
