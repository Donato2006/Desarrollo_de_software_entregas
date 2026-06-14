package services

import (
	"backend/dao"
	"backend/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRegistrarUsuario_UsuarioValido_RetornaCliente(t *testing.T) {
	dao.Connect()

	correo := "test_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := domain.Usuario{
		Nombre:   "Usuario Test",
		Correo:   correo,
		Password: "123456",
	}

	usuarioCreado, err := RegistrarUsuario(usuario)

	assert.NoError(t, err)
	assert.NotZero(t, usuarioCreado.ID)
	assert.Equal(t, "cliente", usuarioCreado.Rol)
	assert.NotEqual(t, "123456", usuarioCreado.Password)

	t.Cleanup(func() {
		dao.DB.Delete(&usuarioCreado)
	})
}

func TestRegistrarUsuario_CorreoDuplicado_RetornaError(t *testing.T) {
	dao.Connect()

	correo := "duplicado_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := domain.Usuario{
		Nombre:   "Usuario Test",
		Correo:   correo,
		Password: "123456",
	}

	usuarioCreado, err := RegistrarUsuario(usuario)
	assert.NoError(t, err)

	_, err = RegistrarUsuario(usuario)

	assert.Error(t, err)
	assert.Equal(t, "el correo ya está registrado", err.Error())

	t.Cleanup(func() {
		dao.DB.Delete(&usuarioCreado)
	})
}

func TestLogin_CredencialesValidas_RetornaTokenYRol(t *testing.T) {
	dao.Connect()

	correo := "login_" + time.Now().Format("20060102150405") + "@gmail.com"

	usuario := domain.Usuario{
		Nombre:   "Login Test",
		Correo:   correo,
		Password: "123456",
	}

	usuarioCreado, err := RegistrarUsuario(usuario)
	assert.NoError(t, err)

	token, rol, err := Login(correo, "123456")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, "cliente", rol)

	t.Cleanup(func() {
		dao.DB.Delete(&usuarioCreado)
	})
}

func TestLogin_CredencialesInvalidas_RetornaError(t *testing.T) {
	dao.Connect()

	_, _, err := Login("usuario_inexistente@gmail.com", "mal")

	assert.Error(t, err)
	assert.Equal(t, "credenciales inválidas", err.Error())
}
