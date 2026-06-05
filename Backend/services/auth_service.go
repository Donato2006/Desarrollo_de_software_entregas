package services

import (
	"backend/dao"
	"backend/domain"
	"backend/utils"
	"errors"
)

func RegistrarUsuario(usuario domain.Usuario) (domain.Usuario, error) {
	var existente domain.Usuario

	resultado := dao.DB.Where("correo = ?", usuario.Correo).First(&existente)

	if resultado.Error == nil {
		return usuario, errors.New("el correo ya está registrado")
	}

	passwordHasheada, err := utils.HashPassword(usuario.Password)

	if err != nil {
		return usuario, err
	}

	usuario.Password = passwordHasheada

	if usuario.Rol == "" {
		usuario.Rol = "cliente"
	}

	resultado = dao.DB.Create(&usuario)

	return usuario, resultado.Error
}

func Login(correo string, password string) (string, string, error) {

	var usuario domain.Usuario

	resultado := dao.DB.Where("correo = ?", correo).First(&usuario)

	if resultado.Error != nil {
		return "", "", errors.New("credenciales inválidas")
	}

	passwordCorrecta := utils.CheckPassword(password, usuario.Password)

	if !passwordCorrecta {
		return "", "", errors.New("credenciales inválidas")
	}

	token, err := utils.GenerarJWT(usuario.ID, usuario.Rol)

	if err != nil {
		return "", "", err
	}

	return token, usuario.Rol, nil
}
