package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"praisindo/handler"
	"testing"
)

func TestGetHelloMessage(t *testing.T) {
	t.Run("Positive Case - Correct Message", func(t *testing.T) {
		expectedOutput := "Halo dari Gin!"
		actualOutput := handler.GetHelloMessage()
		require.Equal(t, expectedOutput, actualOutput, "The message should be '%s'", expectedOutput)
	})
}

func TestRootHandler(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup the router and route
	router := gin.Default()
	router.GET("/", handler.RootHandler)

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/", nil)

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code) // Check if the status code is 200

	expectedBody := `{"message":"Halo dari Gin!"}`
	assert.JSONEq(t, expectedBody, w.Body.String()) // Check if the response body matches the expected JSON
}

// Struct untuk menampung data JSON dalam tes
type JsonRequest struct {
	Message string `json:"message"`
}

func TestPostHandler(t *testing.T) {
	// Setup router
	r := gin.Default()
	r.POST("/", handler.PostHandler)

	t.Run("Positive Case", func(t *testing.T) {
		// Persiapan data JSON
		requestBody := JsonRequest{Message: "Hello from test!"}
		requestBodyBytes, _ := json.Marshal(requestBody)

		// Buat permintaan HTTP POST
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Buat ResponseRecorder untuk merekam respons
		w := httptest.NewRecorder()

		// Lakukan permintaan
		r.ServeHTTP(w, req)

		// Periksa status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Periksa body respons
		expectedBody := `{"message":"Hello from test!"}`
		assert.JSONEq(t, expectedBody, w.Body.String())
	})

	t.Run("Negative Case - EOF Error", func(t *testing.T) {
		// Persiapan data JSON yang salah
		requestBody := ""
		requestBodyBytes := []byte(requestBody)
		// Buat permintaan HTTP POST
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Content-Type", "application/json")
		// Buat ResponseRecorder untuk merekam respons
		w := httptest.NewRecorder()
		// Lakukan permintaan
		r.ServeHTTP(w, req)
		// Periksa status code
		assert.Equal(t, http.StatusBadRequest, w.Code)
		// Periksa body respons
		assert.Contains(t, w.Body.String(), "{\"error\":\"EOF\"}")
	})
}
