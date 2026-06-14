package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func generarTokenTest(usuarioID uint, rol string) string {
	claims := jwt.MapClaims{
		"usuario_id": usuarioID,
		"rol":        rol,
		"exp":        time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(jwtSecret)

	return tokenString
}

func routerTestMiddleware() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.GET("/protegida", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "ok",
		})
	})

	r.GET("/admin", AuthMiddleware(), AdminMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "admin ok",
		})
	})

	return r
}

func TestAuthMiddleware_SinToken_Retorna401(t *testing.T) {
	r := routerTestMiddleware()

	req := httptest.NewRequest(http.MethodGet, "/protegida", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_TokenValido_Retorna200(t *testing.T) {
	r := routerTestMiddleware()

	token := generarTokenTest(1, "cliente")

	req := httptest.NewRequest(http.MethodGet, "/protegida", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAdminMiddleware_UsuarioCliente_Retorna403(t *testing.T) {
	r := routerTestMiddleware()

	token := generarTokenTest(1, "cliente")

	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestAdminMiddleware_UsuarioAdmin_Retorna200(t *testing.T) {
	r := routerTestMiddleware()

	token := generarTokenTest(1, "admin")

	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthMiddleware_TokenInvalido_Retorna401(t *testing.T) {
	r := routerTestMiddleware()

	req := httptest.NewRequest(http.MethodGet, "/protegida", nil)
	req.Header.Set("Authorization", "Bearer token_invalido")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAdminMiddleware_SinRol_Retorna401(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.GET("/admin-sin-rol", AdminMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "admin ok",
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/admin-sin-rol", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
