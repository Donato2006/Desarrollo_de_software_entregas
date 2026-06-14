package controllers

import (
	"backend/dao"
	"backend/domain"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupEntradaRouterTest() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.POST("/entradas", func(c *gin.Context) {
		c.Set("usuario_id", float64(1))
		ComprarEntrada(c)
	})

	r.GET("/mis-entradas", func(c *gin.Context) {
		c.Set("usuario_id", float64(1))
		MisEntradas(c)
	})

	r.DELETE("/entradas/:id", func(c *gin.Context) {
		c.Set("usuario_id", float64(1))
		CancelarEntrada(c)
	})

	r.PUT("/entradas/:id/transferir", func(c *gin.Context) {
		c.Set("usuario_id", float64(1))
		TransferirEntrada(c)
	})

	return r
}

func TestComprarEntrada_BodyInvalido_Retorna400(t *testing.T) {
	r := setupEntradaRouterTest()

	req := httptest.NewRequest(
		http.MethodPost,
		"/entradas",
		bytes.NewReader([]byte("{")),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestComprarEntrada_ConciertoInexistente_Retorna400(t *testing.T) {
	dao.Connect()

	r := setupEntradaRouterTest()

	body := map[string]uint{
		"ConciertoID": 999999,
	}

	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/entradas",
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMisEntradas_Retorna200(t *testing.T) {
	dao.Connect()

	r := setupEntradaRouterTest()

	req := httptest.NewRequest(http.MethodGet, "/mis-entradas", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCancelarEntrada_Inexistente_Retorna400(t *testing.T) {
	dao.Connect()

	r := setupEntradaRouterTest()

	req := httptest.NewRequest(http.MethodDelete, "/entradas/999999", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTransferirEntrada_BodyInvalido_Retorna400(t *testing.T) {
	r := setupEntradaRouterTest()

	req := httptest.NewRequest(
		http.MethodPut,
		"/entradas/1/transferir",
		bytes.NewReader([]byte("{")),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTransferirEntrada_Inexistente_Retorna400(t *testing.T) {
	dao.Connect()

	r := setupEntradaRouterTest()

	body := map[string]string{
		"CorreoDestino": "noexiste@test.com",
	}

	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPut,
		"/entradas/999999/transferir",
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestComprarEntrada_Valida_Retorna201(t *testing.T) {
	dao.Connect()

	r := gin.New()

	usuario := domain.Usuario{
		Nombre:   "Usuario Entrada Controller",
		Correo:   "entrada_controller_valida@test.com",
		Password: "123456",
		Rol:      "cliente",
	}

	dao.DB.Create(&usuario)

	concierto := domain.Concierto{
		Nombre:           "Concierto Entrada Controller",
		Fecha:            time.Now().AddDate(0, 1, 0),
		CupoTotal:        5,
		CuposDisponibles: 5,
		Lugar:            "Lugar Test",
	}

	dao.DB.Create(&concierto)

	r.POST("/entradas", func(c *gin.Context) {
		c.Set("usuario_id", float64(usuario.ID))
		ComprarEntrada(c)
	})

	body := map[string]uint{
		"ConciertoID": concierto.ID,
	}

	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/entradas",
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	t.Cleanup(func() {
		dao.DB.Where("usuario_id = ?", usuario.ID).Delete(&domain.Entrada{})
		dao.DB.Delete(&usuario)
		dao.DB.Delete(&concierto)
	})
}

func TestCancelarEntrada_Valida_Retorna200(t *testing.T) {
	dao.Connect()

	r := gin.New()

	usuario := domain.Usuario{
		Nombre:   "Usuario Cancelar Controller",
		Correo:   "cancelar_controller@test.com",
		Password: "123456",
		Rol:      "cliente",
	}

	dao.DB.Create(&usuario)

	concierto := domain.Concierto{
		Nombre:           "Concierto Cancelar Controller",
		Fecha:            time.Now().AddDate(0, 1, 0),
		CupoTotal:        5,
		CuposDisponibles: 5,
		Lugar:            "Lugar Test",
	}

	dao.DB.Create(&concierto)

	entrada := domain.Entrada{
		UsuarioID:   usuario.ID,
		ConciertoID: concierto.ID,
		Estado:      "activa",
	}

	dao.DB.Create(&entrada)

	r.DELETE("/entradas/:id", func(c *gin.Context) {
		c.Set("usuario_id", float64(usuario.ID))
		CancelarEntrada(c)
	})

	req := httptest.NewRequest(
		http.MethodDelete,
		"/entradas/"+strconv.Itoa(int(entrada.ID)),
		nil,
	)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	t.Cleanup(func() {
		dao.DB.Delete(&entrada)
		dao.DB.Delete(&usuario)
		dao.DB.Delete(&concierto)
	})
}
