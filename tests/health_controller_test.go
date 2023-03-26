package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohidex/identity-service/controllers"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckRoute(t *testing.T) {
	router := SetUpRouter()
	healthController := new(controllers.HealthController)
	router.GET("/health", healthController.Status)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Working!", w.Body.String())
}
