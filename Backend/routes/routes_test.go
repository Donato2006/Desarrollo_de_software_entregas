package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes_LoginExiste_Retorna400ConBodyInvalido(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	SetupRoutes(r)

	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		strings.NewReader("{"),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSetupRoutes_RegisterExiste_Retorna400ConBodyInvalido(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	SetupRoutes(r)

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		strings.NewReader("{"),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSetupRoutes_EntradasProtegidaSinToken_Retorna401(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	SetupRoutes(r)

	req := httptest.NewRequest(http.MethodPost, "/entradas", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
