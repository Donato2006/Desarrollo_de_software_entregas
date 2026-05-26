package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/conciertos", controllers.ObtenerConciertos)

	r.POST("/conciertos", middleware.AuthMiddleware(), controllers.CrearConcierto)

	r.GET("/conciertos/:id", controllers.ObtenerConciertoPorID)

	r.PUT("/conciertos/:id", middleware.AuthMiddleware(), controllers.ActualizarConcierto)

	r.DELETE("/conciertos/:id", middleware.AuthMiddleware(), controllers.EliminarConcierto)

	r.POST("/register", controllers.Register)

	r.POST("/login", controllers.Login)

	r.POST("/entradas", middleware.AuthMiddleware(), controllers.ComprarEntrada)
}
