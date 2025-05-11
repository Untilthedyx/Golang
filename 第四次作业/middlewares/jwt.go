package middlewares

import (
	"Four/config"
	"Four/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer")
		if tokenString == "" {
			c.AbortWithStatusJSON(utils.TokenError.Code, gin.H{"error": utils.TokenError.Msg})
			return
		}
		if claims, err := utils.ParseToken(tokenString, config.JWTSecret); err != nil {
			c.AbortWithStatusJSON(utils.TokenError.Code, gin.H{"error": utils.TokenError.Msg})
			return
		} else {
			c.Set("user_id", claims.Id)
			c.Set("username", claims.Username)
			c.Set("email", claims.Email)
			c.Next()
		}

	}
}
