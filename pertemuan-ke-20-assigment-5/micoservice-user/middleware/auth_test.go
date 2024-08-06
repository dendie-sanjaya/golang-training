package middleware_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"praisindo/middleware"
	"testing"
)

func TestAuthMiddleware_PositiveCase(t *testing.T) {
	// Inisialisasi router
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	// Handler yang hanya dapat diakses dengan token valid
	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Private data"})
	})

	// Buat permintaan HTTP GET ke endpoint "/private" dengan token valid
	req, _ := http.NewRequest("GET", "/private", nil)
	req.Header.Set("Authorization", "valid-token")
	w := httptest.NewRecorder()

	// Lakukan permintaan
	r.ServeHTTP(w, req)

	// Periksa status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Periksa body respons
	assert.JSONEq(t, `{"message":"Private data"}`, w.Body.String())
}

func TestAuthMiddleware_NegativeCase_NoToken(t *testing.T) {
	// Inisialisasi router
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	// Handler yang hanya dapat diakses dengan token valid
	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Private data"})
	})

	// Buat permintaan HTTP GET ke endpoint "/private" tanpa token
	req, _ := http.NewRequest("GET", "/private", nil)
	w := httptest.NewRecorder()

	// Lakukan permintaan
	r.ServeHTTP(w, req)

	// Periksa status code
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Periksa body respons
	assert.Contains(t, w.Body.String(), "Authorization token required")
}

func TestAuthMiddleware_NegativeCase_InvalidToken(t *testing.T) {
	// Inisialisasi router
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	// Handler yang hanya dapat diakses dengan token valid
	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Private data"})
	})

	// Buat permintaan HTTP GET ke endpoint "/private" dengan token tidak valid
	req, _ := http.NewRequest("GET", "/private", nil)
	req.Header.Set("Authorization", "invalid-token")
	w := httptest.NewRecorder()

	// Lakukan permintaan
	r.ServeHTTP(w, req)

	// Periksa status code
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Periksa body respons
	assert.Contains(t, w.Body.String(), "Invalid authorization token")
}
