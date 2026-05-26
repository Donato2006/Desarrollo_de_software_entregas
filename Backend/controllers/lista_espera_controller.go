package controllers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnotarseListaEspera(c *gin.Context) {

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

	lista, err := services.AnotarseListaEspera(usuarioID, datos.ConciertoID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, lista)
}

func MisListasEspera(c *gin.Context) {

	usuarioIDFloat := c.GetFloat64("usuario_id")
	usuarioID := uint(usuarioIDFloat)

	listas, err := services.ObtenerListasEsperaUsuario(usuarioID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener listas de espera",
		})
		return
	}

	c.JSON(http.StatusOK, listas)
}

func SalirListaEspera(c *gin.Context) {

	id := c.Param("id")

	usuarioIDFloat := c.GetFloat64("usuario_id")
	usuarioID := uint(usuarioIDFloat)

	err := services.SalirListaEspera(usuarioID, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Saliste de la lista de espera correctamente",
	})
}

func VerListaEsperaConcierto(c *gin.Context) {
	conciertoID := c.Param("conciertoId")

	listas, err := services.ObtenerListaEsperaConcierto(conciertoID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener lista de espera",
		})
		return
	}

	c.JSON(http.StatusOK, listas)
}
