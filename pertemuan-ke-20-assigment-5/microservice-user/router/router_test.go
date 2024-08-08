// router/router_test.go
package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"praisindo/router"
)

func TestPrivateRoutes(t *testing.T) {
	// Set gin ke mode test
	gin.SetMode(gin.TestMode)
	// Buat router gin baru
	r := gin.Default()
	// Buat mock user handler
	mockUserHandler := &handler.MockUserHandler{}
	// Set up router dengan mock handler
	router.SetupRouter(r, mockUserHandler)

	// Fungsi helper untuk menambahkan header autentikasi dasar
	addAuth := func(req *http.Request) {
		req.SetBasicAuth("user", "pass")
	}

	// Test POST /users
	t.Run("POST /users", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/users", nil)
		addAuth(req)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Verifikasi bahwa status code adalah 201 Created
		require.Equal(t, http.StatusCreated, w.Code)
		// Verifikasi bahwa body mengandung "user created"
		require.Contains(t, w.Body.String(), "user created")
	})

	// Test PUT /users/:id
	t.Run("PUT /users/:id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/users/1", nil)
		addAuth(req)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Verifikasi bahwa status code adalah 200 OK
		require.Equal(t, http.StatusOK, w.Code)
		// Verifikasi bahwa body mengandung "user updated"
		require.Contains(t, w.Body.String(), "user updated")
	})

	// Test DELETE /users/:id
	t.Run("DELETE /users/:id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
		addAuth(req)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Verifikasi bahwa status code adalah 200 OK
		require.Equal(t, http.StatusOK, w.Code)
		// Verifikasi bahwa body mengandung "user deleted"
		require.Contains(t, w.Body.String(), "user deleted")
	})
}

func TestPrivateRoutesUnauthorized(t *testing.T) {
	// Set gin ke mode test
	gin.SetMode(gin.TestMode)
	// Buat router gin baru
	r := gin.Default()
	// Buat mock user handler
	mockUserHandler := &handler.MockUserHandler{}
	// Set up router dengan mock handler
	router.SetupRouter(r, mockUserHandler)

	// Test POST /users tanpa autentikasi
	t.Run("POST /users - Unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/users", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Verifikasi bahwa status code adalah 401 Unauthorized
		require.Equal(t, http.StatusUnauthorized, w.Code)
		// Verifikasi bahwa body mengandung "Authorization basic token required"
		require.Contains(t, w.Body.String(), "Authorization basic token required")
	})

	// Test PUT /users/:id tanpa autentikasi
	t.Run("PUT /users/:id - Unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/users/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Verifikasi bahwa status code adalah 401 Unauthorized
		require.Equal(t, http.StatusUnauthorized, w.Code)
		// Verifikasi bahwa body mengandung "Authorization basic token required"
		require.Contains(t, w.Body.String(), "Authorization basic token required")
	})

	// Test DELETE /users/:id tanpa autentikasi
	t.Run("DELETE /users/:id - Unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Verifikasi bahwa status code adalah 401 Unauthorized
		require.Equal(t, http.StatusUnauthorized, w.Code)
		// Verifikasi bahwa body mengandung "Authorization basic token required"
		require.Contains(t, w.Body.String(), "Authorization basic token required")
	})
}
