package controllers

import (
	"backend/dao"
	"backend/domain"
	"backend/services"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupAuthRouterTest() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.POST("/register", Register)
	r.POST("/login", Login)

	return r
}

func TestRegister_DatosValidos_Retorna201(t *testing.T) {
	dao.Connect()

	r := setupAuthRouterTest()

	correo := "controller_reg_" + time.Now().Format("20060102150405") + "@gmail.com"

	body := map[string]string{
		"Nombre":   "Usuario Controller",
		"Correo":   correo,
		"Password": "123456",
	}

	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var usuario domain.Usuario

	dao.DB.Where("correo = ?", correo).First(&usuario)

	t.Cleanup(func() {
		dao.DB.Delete(&usuario)
	})
}

func TestRegister_BodyInvalido_Retorna400(t *testing.T) {
	r := setupAuthRouterTest()

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewReader([]byte("{")),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogin_CredencialesValidas_Retorna200YToken(t *testing.T) {
	dao.Connect()

	r := setupAuthRouterTest()

	correo := "controller_login_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := domain.Usuario{
		Nombre:   "Login Controller",
		Correo:   correo,
		Password: "123456",
	}

	usuarioCreado, err := services.RegistrarUsuario(usuario)

	assert.NoError(t, err)

	body := map[string]string{
		"Correo":   correo,
		"Password": "123456",
	}

	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
	assert.Contains(t, w.Body.String(), "rol")

	t.Cleanup(func() {
		dao.DB.Delete(&usuarioCreado)
	})
}

func TestLogin_CredencialesInvalidas_Retorna401(t *testing.T) {
	dao.Connect()

	r := setupAuthRouterTest()

	body := map[string]string{
		"Correo":   "noexiste_controller@gmail.com",
		"Password": "mal",
	}

	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
