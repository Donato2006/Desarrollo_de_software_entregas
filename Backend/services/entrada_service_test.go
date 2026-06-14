package services

import (
	"backend/dao"
	"backend/domain"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func crearUsuarioTest(t *testing.T, correo string) domain.Usuario {
	usuario := domain.Usuario{
		Nombre:   "Usuario Test",
		Correo:   correo,
		Password: "123456",
		Rol:      "cliente",
	}

	usuarioCreado, err := RegistrarUsuario(usuario)

	assert.NoError(t, err)

	t.Cleanup(func() {
		dao.DB.Delete(&usuarioCreado)
	})

	return usuarioCreado
}

func crearConciertoTest(t *testing.T, cupos int) domain.Concierto {
	concierto := domain.Concierto{
		Nombre:           "Concierto Test",
		Fecha:            time.Now().AddDate(0, 1, 0),
		Lugar:            "Lugar Test",
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

func TestComprarEntrada_ConCupos_CreaEntradaYDescuentaCupo(t *testing.T) {
	dao.Connect()

	correo := "compra_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := crearUsuarioTest(t, correo)

	concierto := crearConciertoTest(t, 2)

	entrada, err := ComprarEntrada(usuario.ID, concierto.ID)

	assert.NoError(t, err)
	assert.NotZero(t, entrada.ID)
	assert.Equal(t, usuario.ID, entrada.UsuarioID)
	assert.Equal(t, concierto.ID, entrada.ConciertoID)
	assert.Equal(t, "activa", entrada.Estado)

	var conciertoActualizado domain.Concierto

	dao.DB.First(&conciertoActualizado, concierto.ID)

	assert.Equal(t, 1, conciertoActualizado.CuposDisponibles)

	t.Cleanup(func() {
		dao.DB.Delete(&entrada)
	})
}

func TestComprarEntrada_SinCupos_RetornaError(t *testing.T) {
	dao.Connect()

	correo := "sincupos_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := crearUsuarioTest(t, correo)

	concierto := crearConciertoTest(t, 0)

	_, err := ComprarEntrada(usuario.ID, concierto.ID)

	assert.Error(t, err)
	assert.Equal(t, "no hay cupos disponibles", err.Error())
}

func TestCancelarEntrada_EntradaActiva_CancelaYLiberaCupo(t *testing.T) {
	dao.Connect()

	correo := "cancelar_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := crearUsuarioTest(t, correo)

	concierto := crearConciertoTest(t, 1)

	entrada, err := ComprarEntrada(usuario.ID, concierto.ID)

	assert.NoError(t, err)

	entradaID := strconv.Itoa(int(entrada.ID))

	err = CancelarEntrada(usuario.ID, entradaID)

	assert.NoError(t, err)

	var entradaActualizada domain.Entrada

	dao.DB.First(&entradaActualizada, entrada.ID)

	assert.Equal(t, "cancelada", entradaActualizada.Estado)

	var conciertoActualizado domain.Concierto

	dao.DB.First(&conciertoActualizado, concierto.ID)

	assert.Equal(t, 1, conciertoActualizado.CuposDisponibles)

	t.Cleanup(func() {
		dao.DB.Delete(&entradaActualizada)
	})
}

func TestTransferirEntrada_EntradaActiva_CambiaUsuario(t *testing.T) {
	dao.Connect()

	correoOrigen := "origen_" + time.Now().Format("20060102150405") + "@gmail.com"

	correoDestino := "destino_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuarioOrigen := crearUsuarioTest(t, correoOrigen)

	usuarioDestino := crearUsuarioTest(t, correoDestino)

	concierto := crearConciertoTest(t, 1)

	entrada, err := ComprarEntrada(usuarioOrigen.ID, concierto.ID)

	assert.NoError(t, err)

	entradaID := strconv.Itoa(int(entrada.ID))

	err = TransferirEntrada(usuarioOrigen.ID, entradaID, usuarioDestino.Correo)

	assert.NoError(t, err)

	var entradaActualizada domain.Entrada

	dao.DB.First(&entradaActualizada, entrada.ID)

	assert.Equal(t, usuarioDestino.ID, entradaActualizada.UsuarioID)

	t.Cleanup(func() {
		dao.DB.Delete(&entradaActualizada)
	})
}

func TestTransferirEntrada_UsuarioDestinoInexistente_RetornaError(t *testing.T) {
	dao.Connect()

	correoOrigen := "origen_error_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuarioOrigen := crearUsuarioTest(t, correoOrigen)

	concierto := crearConciertoTest(t, 1)

	entrada, err := ComprarEntrada(usuarioOrigen.ID, concierto.ID)

	assert.NoError(t, err)

	entradaID := strconv.Itoa(int(entrada.ID))

	err = TransferirEntrada(usuarioOrigen.ID, entradaID, "noexiste@gmail.com")

	assert.Error(t, err)
	assert.Equal(t, "usuario destino no encontrado", err.Error())

	t.Cleanup(func() {
		dao.DB.Delete(&entrada)
	})
}

func TestObtenerEntradasUsuario_RetornaEntradasDelUsuario(t *testing.T) {
	dao.Connect()

	correo := "obtener_entradas_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := crearUsuarioTest(t, correo)

	concierto := crearConciertoTest(t, 2)

	entrada, err := ComprarEntrada(usuario.ID, concierto.ID)

	assert.NoError(t, err)

	entradas, err := ObtenerEntradasUsuario(usuario.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, entradas)

	t.Cleanup(func() {
		dao.DB.Delete(&entrada)
	})
}
