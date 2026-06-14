package controllers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReporteOcupacion(c *gin.Context) {

	reporte, err := services.ObtenerReporteOcupacion()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo generar el reporte",
		})

		return
	}

	c.JSON(http.StatusOK, reporte)
}
