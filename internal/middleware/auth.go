package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Zekeriyyah/stellar-x/pkg"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// access the authHeader for the token
	
		tokenHeader := c.GetHeader("Authorization")
		token, err := pkg.ExtractTokenStr(tokenHeader)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing token"})
			c.Abort()
			return
		}

		claims, err := pkg.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// log user IP, device info, country, and browser
		

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
