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

	r.GET("/mis-entradas", middleware.AuthMiddleware(), controllers.MisEntradas)

	r.DELETE("/entradas/:id", middleware.AuthMiddleware(), controllers.CancelarEntrada)

	r.POST("/lista-espera", middleware.AuthMiddleware(), controllers.AnotarseListaEspera)

	r.GET("/mis-listas-espera", middleware.AuthMiddleware(), controllers.MisListasEspera)

	r.DELETE("/lista-espera/:id", middleware.AuthMiddleware(), controllers.SalirListaEspera)

	r.GET("/lista-espera/:conciertoId", middleware.AuthMiddleware(), controllers.VerListaEsperaConcierto)
}
