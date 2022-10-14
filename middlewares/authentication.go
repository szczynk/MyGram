package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/szczynk/MyGram/helpers"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		} else {
			c.Set("userData", verifyToken)
			c.Next()
		}
	}
}
