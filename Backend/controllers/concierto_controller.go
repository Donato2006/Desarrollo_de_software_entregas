package controllers

import (
	"backend/database"
	"backend/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ObtenerConciertos(c *gin.Context) {

	var conciertos []domain.Concierto

	database.DB.Find(&conciertos)

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

	database.DB.Create(&concierto)

	c.JSON(http.StatusCreated, concierto)
}

func ObtenerConciertoPorID(c *gin.Context) {
	id := c.Param("id")

	var concierto domain.Concierto

	resultado := database.DB.First(&concierto, id)

	if resultado.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Concierto no encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, concierto)
}
