package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uname, pwd, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau Password tidak boleh kosong"})
			c.Abort() 
			return
		}

		if uname == "admin" && pwd == "admin" {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau Password tidak sesuai"})
		c.Abort()
	}
}
