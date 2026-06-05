package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var usuario domain.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
		})
		return
	}

	usuarioCreado, err := services.RegistrarUsuario(usuario)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, usuarioCreado)
}

func Login(c *gin.Context) {

	var datosLogin struct {
		Correo   string
		Password string
	}

	if err := c.ShouldBindJSON(&datosLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
		})
		return
	}

	token, rol, err := services.Login(
		datosLogin.Correo,
		datosLogin.Password,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"rol":   rol,
	})
}
