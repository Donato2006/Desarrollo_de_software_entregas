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

func setupConciertoRouterTest() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.GET("/conciertos", ObtenerConciertos)
	r.POST("/conciertos", CrearConcierto)
	r.GET("/conciertos/:id", ObtenerConciertoPorID)
	r.PUT("/conciertos/:id", ActualizarConcierto)
	r.DELETE("/conciertos/:id", EliminarConcierto)

	return r
}

func TestObtenerConciertos_Retorna200(t *testing.T) {
	dao.Connect()

	r := setupConciertoRouterTest()

	req := httptest.NewRequest(http.MethodGet, "/conciertos", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCrearConcierto_DatosValidos_Retorna201(t *testing.T) {
	dao.Connect()

	r := setupConciertoRouterTest()

	body := domain.Concierto{
		Nombre:           "Concierto Controller Test",
		Fecha:            time.Now().AddDate(0, 1, 0),
		Lugar:            "Lugar Test",
		CupoTotal:        100,
		CuposDisponibles: 100,
	}

	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/conciertos",
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var concierto domain.Concierto
	dao.DB.Where("nombre = ?", "Concierto Controller Test").First(&concierto)

	t.Cleanup(func() {
		dao.DB.Delete(&concierto)
	})
}

func TestCrearConcierto_BodyInvalido_Retorna400(t *testing.T) {
	r := setupConciertoRouterTest()

	req := httptest.NewRequest(
		http.MethodPost,
		"/conciertos",
		bytes.NewReader([]byte("{")),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestObtenerConciertoPorID_Existente_Retorna200(t *testing.T) {
	dao.Connect()

	r := setupConciertoRouterTest()

	concierto := domain.Concierto{
		Nombre:           "Concierto Buscar Test",
		Fecha:            time.Now().AddDate(0, 1, 0),
		Lugar:            "Lugar Buscar",
		CupoTotal:        50,
		CuposDisponibles: 50,
	}

	dao.DB.Create(&concierto)

	id := strconv.Itoa(int(concierto.ID))

	req := httptest.NewRequest(http.MethodGet, "/conciertos/"+id, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	t.Cleanup(func() {
		dao.DB.Delete(&concierto)
	})
}

func TestObtenerConciertoPorID_Inexistente_Retorna404(t *testing.T) {
	dao.Connect()

	r := setupConciertoRouterTest()

	req := httptest.NewRequest(http.MethodGet, "/conciertos/999999", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestActualizarConcierto_Existente_Retorna200(t *testing.T) {
	dao.Connect()

	r := setupConciertoRouterTest()

	concierto := domain.Concierto{
		Nombre:           "Concierto Antes",
		Fecha:            time.Now().AddDate(0, 1, 0),
		Lugar:            "Lugar Antes",
		CupoTotal:        20,
		CuposDisponibles: 20,
	}

	dao.DB.Create(&concierto)

	datosActualizados := domain.Concierto{
		Nombre:           "Concierto Después",
		Fecha:            time.Now().AddDate(0, 2, 0),
		Lugar:            "Lugar Después",
		CupoTotal:        30,
		CuposDisponibles: 30,
	}

	bodyBytes, _ := json.Marshal(datosActualizados)

	id := strconv.Itoa(int(concierto.ID))

	req := httptest.NewRequest(
		http.MethodPut,
		"/conciertos/"+id,
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	t.Cleanup(func() {
		dao.DB.Delete(&concierto)
	})
}

func TestEliminarConcierto_Existente_Retorna200(t *testing.T) {
	dao.Connect()

	r := setupConciertoRouterTest()

	concierto := domain.Concierto{
		Nombre:           "Concierto Eliminar",
		Fecha:            time.Now().AddDate(0, 1, 0),
		Lugar:            "Lugar Eliminar",
		CupoTotal:        10,
		CuposDisponibles: 10,
	}

	dao.DB.Create(&concierto)

	id := strconv.Itoa(int(concierto.ID))

	req := httptest.NewRequest(http.MethodDelete, "/conciertos/"+id, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEliminarConcierto_Inexistente_Retorna404(t *testing.T) {
	dao.Connect()

	r := setupConciertoRouterTest()

	req := httptest.NewRequest(http.MethodDelete, "/conciertos/999999", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
