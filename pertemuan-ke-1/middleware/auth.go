package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware untuk autentikasi
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// Periksa apakah token disediakan
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Verifikasi token (misalnya, cocokkan dengan token yang diharapkan)
		if token != "valid-token" { // ganti "valid-token" dengan logika validasi token sebenarnya
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		c.Next()
	}
}
