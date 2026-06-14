package services

import (
	"backend/dao"
	"backend/domain"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCrearConcierto_DatosValidos_CreaConcierto(t *testing.T) {
	dao.Connect()

	concierto := domain.Concierto{
		Nombre:           "Service Crear Test",
		Fecha:            time.Now().AddDate(0, 1, 0),
		Lugar:            "Lugar Service",
		CupoTotal:        100,
		CuposDisponibles: 100,
	}

	conciertoCreado, err := CrearConcierto(concierto)

	assert.NoError(t, err)
	assert.NotZero(t, conciertoCreado.ID)
	assert.Equal(t, "Service Crear Test", conciertoCreado.Nombre)

	t.Cleanup(func() {
		dao.DB.Delete(&conciertoCreado)
	})
}

func TestObtenerTodosLosConciertos_RetornaLista(t *testing.T) {
	dao.Connect()

	concierto := crearConciertoTest(t, 5)

	conciertos, err := ObtenerTodosLosConciertos()

	assert.NoError(t, err)
	assert.NotEmpty(t, conciertos)

	t.Cleanup(func() {
		dao.DB.Delete(&concierto)
	})
}

func TestObtenerConciertoPorID_Existente_RetornaConcierto(t *testing.T) {
	dao.Connect()

	concierto := crearConciertoTest(t, 10)

	id := strconv.Itoa(int(concierto.ID))

	conciertoEncontrado, err := ObtenerConciertoPorID(id)

	assert.NoError(t, err)
	assert.Equal(t, concierto.ID, conciertoEncontrado.ID)

	t.Cleanup(func() {
		dao.DB.Delete(&concierto)
	})
}

func TestObtenerConciertoPorID_Inexistente_RetornaError(t *testing.T) {
	dao.Connect()

	_, err := ObtenerConciertoPorID("999999")

	assert.Error(t, err)
}

func TestActualizarConcierto_Existente_ActualizaDatos(t *testing.T) {
	dao.Connect()

	concierto := crearConciertoTest(t, 20)

	id := strconv.Itoa(int(concierto.ID))

	datos := domain.Concierto{
		Nombre:           "Service Actualizado",
		Fecha:            time.Now().AddDate(0, 2, 0),
		Lugar:            "Lugar Actualizado",
		CupoTotal:        30,
		CuposDisponibles: 30,
	}

	conciertoActualizado, err := ActualizarConcierto(id, datos)

	assert.NoError(t, err)
	assert.Equal(t, "Service Actualizado", conciertoActualizado.Nombre)

	t.Cleanup(func() {
		dao.DB.Delete(&conciertoActualizado)
	})
}

func TestActualizarConcierto_Inexistente_RetornaError(t *testing.T) {
	dao.Connect()

	datos := domain.Concierto{
		Nombre:           "No Existe",
		Fecha:            time.Now(),
		Lugar:            "Nada",
		CupoTotal:        1,
		CuposDisponibles: 1,
	}

	_, err := ActualizarConcierto("999999", datos)

	assert.Error(t, err)
}

func TestEliminarConcierto_Existente_EliminaConcierto(t *testing.T) {
	dao.Connect()

	concierto := crearConciertoTest(t, 10)

	id := strconv.Itoa(int(concierto.ID))

	err := EliminarConcierto(id)

	assert.NoError(t, err)
}

func TestEliminarConcierto_Inexistente_RetornaError(t *testing.T) {
	dao.Connect()

	err := EliminarConcierto("999999")

	assert.Error(t, err)
}
