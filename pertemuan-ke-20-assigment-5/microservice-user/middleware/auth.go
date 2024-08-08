package middleware

import (
	"net/http"
	"praisindo/config"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware untuk autentikasi
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		// fmt.Print(username, password, ok)
		//token := c.GetHeader("Authorization")

		// Periksa apakah token disediakan
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Verifikasi token (misalnya, cocokkan dengan token yang diharapkan)
		//isValid := false
		isValid := (username == config.AuthBasicUsername) && (password == config.AuthBasicPassword)
		if !isValid { // ganti "valid-token" dengan logika validasi token sebenarnya
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		c.Next()
	}
}
