package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ObtenerConciertos(c *gin.Context) {
	conciertos, err := services.ObtenerTodosLosConciertos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener conciertos",
		})
		return
	}

	c.JSON(http.StatusOK, conciertos)
}

func CrearConcierto(c *gin.Context) {
	var concierto domain.Concierto

	if err := c.ShouldBindJSON(&concierto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
		})
		return
	}

	conciertoCreado, err := services.CrearConcierto(concierto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al crear concierto",
		})
		return
	}

	c.JSON(http.StatusCreated, conciertoCreado)
}

func ObtenerConciertoPorID(c *gin.Context) {
	id := c.Param("id")

	concierto, err := services.ObtenerConciertoPorID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Concierto no encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, concierto)
}

func ActualizarConcierto(c *gin.Context) {
	id := c.Param("id")

	var datosActualizados domain.Concierto

	if err := c.ShouldBindJSON(&datosActualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
		})
		return
	}

	conciertoActualizado, err := services.ActualizarConcierto(id, datosActualizados)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Concierto no encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, conciertoActualizado)
}

func EliminarConcierto(c *gin.Context) {
	id := c.Param("id")

	err := services.EliminarConcierto(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Concierto no encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Concierto eliminado correctamente",
	})
}
