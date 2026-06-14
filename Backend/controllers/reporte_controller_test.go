package controllers

import (
	"backend/dao"
	"backend/domain"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestReporteOcupacion_Retorna200(t *testing.T) {
	dao.Connect()

	concierto := domain.Concierto{
		Nombre:           "Reporte Controller Test",
		Fecha:            time.Now().AddDate(0, 1, 0),
		Lugar:            "Lugar Reporte",
		CupoTotal:        100,
		CuposDisponibles: 40,
	}

	dao.DB.Create(&concierto)

	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET("/reporte-ocupacion", ReporteOcupacion)

	req := httptest.NewRequest(http.MethodGet, "/reporte-ocupacion", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	t.Cleanup(func() {
		dao.DB.Delete(&concierto)
	})
}
