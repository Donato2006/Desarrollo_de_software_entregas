package controllers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ComprarEntrada(c *gin.Context) {
	var datos struct {
		ConciertoID uint
	}

	if err := c.ShouldBindJSON(&datos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
		})
		return
	}

	usuarioIDFloat := c.GetFloat64("usuario_id")
	usuarioID := uint(usuarioIDFloat)

	entrada, err := services.ComprarEntrada(usuarioID, datos.ConciertoID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, entrada)
}

func MisEntradas(c *gin.Context) {

	usuarioIDFloat := c.GetFloat64("usuario_id")

	usuarioID := uint(usuarioIDFloat)

	entradas, err := services.ObtenerEntradasUsuario(usuarioID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener entradas",
		})
		return
	}

	c.JSON(http.StatusOK, entradas)
}

func CancelarEntrada(c *gin.Context) {
	id := c.Param("id")

	usuarioIDFloat := c.GetFloat64("usuario_id")
	usuarioID := uint(usuarioIDFloat)

	err := services.CancelarEntrada(usuarioID, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Entrada cancelada correctamente",
	})
}
