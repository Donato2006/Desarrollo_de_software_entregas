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
