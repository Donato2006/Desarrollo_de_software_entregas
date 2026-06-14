package services

import (
	"backend/dao"
	"backend/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestObtenerReporteOcupacion_RetornaReporte(t *testing.T) {
	dao.Connect()

	concierto := domain.Concierto{
		Nombre:           "Reporte Test",
		Fecha:            time.Now().AddDate(0, 1, 0),
		Lugar:            "Lugar Reporte",
		CupoTotal:        100,
		CuposDisponibles: 60,
	}

	dao.DB.Create(&concierto)

	reporte, err := ObtenerReporteOcupacion()

	assert.NoError(t, err)
	assert.NotEmpty(t, reporte)

	t.Cleanup(func() {
		dao.DB.Delete(&concierto)
	})
}
