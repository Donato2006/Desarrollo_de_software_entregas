package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/conciertos", controllers.ObtenerConciertos)

	r.POST("/conciertos", controllers.CrearConcierto)

	r.GET("/conciertos/:id", controllers.ObtenerConciertoPorID)

	r.PUT("/conciertos/:id", controllers.ActualizarConcierto)
}
