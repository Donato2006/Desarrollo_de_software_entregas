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

func setupListaEsperaRouterTest(usuarioID uint) *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.POST("/lista-espera", func(c *gin.Context) {
		c.Set("usuario_id", float64(usuarioID))
		AnotarseListaEspera(c)
	})

	r.GET("/mis-listas-espera", func(c *gin.Context) {
		c.Set("usuario_id", float64(usuarioID))
		MisListasEspera(c)
	})

	r.DELETE("/lista-espera/:id", func(c *gin.Context) {
		c.Set("usuario_id", float64(usuarioID))
		SalirListaEspera(c)
	})

	r.GET("/lista-espera/:conciertoId", VerListaEsperaConcierto)

	return r
}

func crearUsuarioListaController(t *testing.T) domain.Usuario {
	usuario := domain.Usuario{
		Nombre:   "Usuario Lista Controller",
		Correo:   "lista_controller_" + time.Now().Format("20060102150405") + "@test.com",
		Password: "123456",
		Rol:      "cliente",
	}

	err := dao.DB.Create(&usuario).Error
	assert.NoError(t, err)

	t.Cleanup(func() {
		dao.DB.Delete(&usuario)
	})

	return usuario
}

func crearConciertoListaController(t *testing.T, cupos int) domain.Concierto {
	concierto := domain.Concierto{
		Nombre:           "Concierto Lista Controller",
		Fecha:            time.Now().AddDate(0, 1, 0),
		Lugar:            "Lugar Lista",
		CupoTotal:        cupos,
		CuposDisponibles: cupos,
	}

	err := dao.DB.Create(&concierto).Error
	assert.NoError(t, err)

	t.Cleanup(func() {
		dao.DB.Delete(&concierto)
	})

	return concierto
}

func TestAnotarseListaEspera_BodyInvalido_Retorna400(t *testing.T) {
	dao.Connect()

	usuario := crearUsuarioListaController(t)
	r := setupListaEsperaRouterTest(usuario.ID)

	req := httptest.NewRequest(
		http.MethodPost,
		"/lista-espera",
		bytes.NewReader([]byte("{")),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAnotarseListaEspera_ConciertoSinCupos_Retorna201(t *testing.T) {
	dao.Connect()

	usuario := crearUsuarioListaController(t)
	concierto := crearConciertoListaController(t, 0)

	r := setupListaEsperaRouterTest(usuario.ID)

	body := map[string]uint{
		"ConciertoID": concierto.ID,
	}

	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/lista-espera",
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var lista domain.ListaEspera
	dao.DB.Where("usuario_id = ? AND concierto_id = ?", usuario.ID, concierto.ID).First(&lista)

	t.Cleanup(func() {
		dao.DB.Delete(&lista)
	})
}

func TestAnotarseListaEspera_ConciertoConCupos_Retorna400(t *testing.T) {
	dao.Connect()

	usuario := crearUsuarioListaController(t)
	concierto := crearConciertoListaController(t, 5)

	r := setupListaEsperaRouterTest(usuario.ID)

	body := map[string]uint{
		"ConciertoID": concierto.ID,
	}

	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/lista-espera",
		bytes.NewReader(bodyBytes),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMisListasEspera_Retorna200(t *testing.T) {
	dao.Connect()

	usuario := crearUsuarioListaController(t)
	r := setupListaEsperaRouterTest(usuario.ID)

	req := httptest.NewRequest(http.MethodGet, "/mis-listas-espera", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSalirListaEspera_Inexistente_Retorna400(t *testing.T) {
	dao.Connect()

	usuario := crearUsuarioListaController(t)
	r := setupListaEsperaRouterTest(usuario.ID)

	req := httptest.NewRequest(http.MethodDelete, "/lista-espera/999999", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSalirListaEspera_RegistroPropio_Retorna200(t *testing.T) {
	dao.Connect()

	usuario := crearUsuarioListaController(t)
	concierto := crearConciertoListaController(t, 0)

	lista := domain.ListaEspera{
		UsuarioID:    usuario.ID,
		ConciertoID:  concierto.ID,
		PosicionCola: 1,
		Estado:       "esperando",
		FechaAlta:    time.Now(),
	}

	dao.DB.Create(&lista)

	r := setupListaEsperaRouterTest(usuario.ID)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/lista-espera/"+strconv.Itoa(int(lista.ID)),
		nil,
	)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	t.Cleanup(func() {
		dao.DB.Delete(&lista)
	})
}

func TestVerListaEsperaConcierto_Retorna200(t *testing.T) {
	dao.Connect()

	usuario := crearUsuarioListaController(t)
	concierto := crearConciertoListaController(t, 0)

	lista := domain.ListaEspera{
		UsuarioID:    usuario.ID,
		ConciertoID:  concierto.ID,
		PosicionCola: 1,
		Estado:       "esperando",
		FechaAlta:    time.Now(),
	}

	dao.DB.Create(&lista)

	r := setupListaEsperaRouterTest(usuario.ID)

	req := httptest.NewRequest(
		http.MethodGet,
		"/lista-espera/"+strconv.Itoa(int(concierto.ID)),
		nil,
	)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	t.Cleanup(func() {
		dao.DB.Delete(&lista)
	})
}
