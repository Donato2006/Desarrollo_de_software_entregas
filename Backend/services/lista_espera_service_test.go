package services

import (
	"backend/dao"
	"backend/domain"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAnotarseListaEspera_ConciertoSinCupos_CreaRegistro(t *testing.T) {
	dao.Connect()

	correo := "lista_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := crearUsuarioTest(t, correo)

	concierto := crearConciertoTest(t, 0)

	lista, err := AnotarseListaEspera(usuario.ID, concierto.ID)

	assert.NoError(t, err)
	assert.NotZero(t, lista.ID)
	assert.Equal(t, usuario.ID, lista.UsuarioID)
	assert.Equal(t, concierto.ID, lista.ConciertoID)
	assert.Equal(t, 1, lista.PosicionCola)
	assert.Equal(t, "esperando", lista.Estado)

	t.Cleanup(func() {
		dao.DB.Delete(&lista)
	})
}

func TestAnotarseListaEspera_ConciertoConCupos_RetornaError(t *testing.T) {
	dao.Connect()

	correo := "lista_cupos_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := crearUsuarioTest(t, correo)

	concierto := crearConciertoTest(t, 3)

	_, err := AnotarseListaEspera(usuario.ID, concierto.ID)

	assert.Error(t, err)
	assert.Equal(t, "todavía hay cupos disponibles, no hace falta lista de espera", err.Error())
}

func TestAnotarseListaEspera_Duplicado_RetornaError(t *testing.T) {
	dao.Connect()

	correo := "lista_dup_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := crearUsuarioTest(t, correo)

	concierto := crearConciertoTest(t, 0)

	lista, err := AnotarseListaEspera(usuario.ID, concierto.ID)
	assert.NoError(t, err)

	_, err = AnotarseListaEspera(usuario.ID, concierto.ID)

	assert.Error(t, err)
	assert.Equal(t, "ya estás anotado en la lista de espera", err.Error())

	t.Cleanup(func() {
		dao.DB.Delete(&lista)
	})
}

func TestObtenerListasEsperaUsuario_RetornaListasDelUsuario(t *testing.T) {
	dao.Connect()

	correo := "mislistas_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := crearUsuarioTest(t, correo)

	concierto := crearConciertoTest(t, 0)

	lista, err := AnotarseListaEspera(usuario.ID, concierto.ID)
	assert.NoError(t, err)

	listas, err := ObtenerListasEsperaUsuario(usuario.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, listas)

	t.Cleanup(func() {
		dao.DB.Delete(&lista)
	})
}

func TestSalirListaEspera_RegistroPropio_EliminaYReordena(t *testing.T) {
	dao.Connect()

	correo1 := "cola1_" + time.Now().Format("20060102150405") + "@gmail.com"
	correo2 := "cola2_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario1 := crearUsuarioTest(t, correo1)
	usuario2 := crearUsuarioTest(t, correo2)

	concierto := crearConciertoTest(t, 0)

	lista1, err := AnotarseListaEspera(usuario1.ID, concierto.ID)
	assert.NoError(t, err)

	lista2, err := AnotarseListaEspera(usuario2.ID, concierto.ID)
	assert.NoError(t, err)

	listaID := strconv.Itoa(int(lista1.ID))

	err = SalirListaEspera(usuario1.ID, listaID)

	assert.NoError(t, err)

	var lista2Actualizada domain.ListaEspera
	dao.DB.First(&lista2Actualizada, lista2.ID)

	assert.Equal(t, 1, lista2Actualizada.PosicionCola)

	t.Cleanup(func() {
		dao.DB.Delete(&lista1)
		dao.DB.Delete(&lista2Actualizada)
	})
}

func TestSalirListaEspera_RegistroAjeno_RetornaError(t *testing.T) {
	dao.Connect()

	correo1 := "ajeno1_" + time.Now().Format("20060102150405") + "@gmail.com"
	correo2 := "ajeno2_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario1 := crearUsuarioTest(t, correo1)
	usuario2 := crearUsuarioTest(t, correo2)

	concierto := crearConciertoTest(t, 0)

	lista, err := AnotarseListaEspera(usuario1.ID, concierto.ID)
	assert.NoError(t, err)

	listaID := strconv.Itoa(int(lista.ID))

	err = SalirListaEspera(usuario2.ID, listaID)

	assert.Error(t, err)
	assert.Equal(t, "no podés eliminar una lista de espera que no es tuya", err.Error())

	t.Cleanup(func() {
		dao.DB.Delete(&lista)
	})
}

func TestObtenerListaEsperaConcierto_RetornaOrdenada(t *testing.T) {
	dao.Connect()

	correo1 := "orden1_" + time.Now().Format("20060102150405") + "@gmail.com"
	correo2 := "orden2_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario1 := crearUsuarioTest(t, correo1)
	usuario2 := crearUsuarioTest(t, correo2)

	concierto := crearConciertoTest(t, 0)

	lista1, err := AnotarseListaEspera(usuario1.ID, concierto.ID)
	assert.NoError(t, err)

	lista2, err := AnotarseListaEspera(usuario2.ID, concierto.ID)
	assert.NoError(t, err)

	conciertoID := strconv.Itoa(int(concierto.ID))

	listas, err := ObtenerListaEsperaConcierto(conciertoID)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(listas), 2)
	assert.Equal(t, 1, listas[0].PosicionCola)
	assert.Equal(t, 2, listas[1].PosicionCola)

	t.Cleanup(func() {
		dao.DB.Delete(&lista1)
		dao.DB.Delete(&lista2)
	})
}
