package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura las rutas REST de la API y las vincula con el motor de Gin
func SetupRoutes(r *gin.Engine) {
	// GET /conciertos -> Devuelve la lista de todos los conciertos
	r.GET("/conciertos", controllers.ObtenerConciertos)

	r.POST("/conciertos", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.CrearConcierto)

	r.GET("/conciertos/:id", controllers.ObtenerConciertoPorID)

	r.PUT("/conciertos/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.ActualizarConcierto)

	r.DELETE("/conciertos/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.EliminarConcierto)

	r.POST("/register", controllers.Register)

	r.POST("/login", controllers.Login)

	r.POST("/entradas", middleware.AuthMiddleware(), controllers.ComprarEntrada)

	r.GET("/mis-entradas", middleware.AuthMiddleware(), controllers.MisEntradas)

	r.DELETE("/entradas/:id", middleware.AuthMiddleware(), controllers.CancelarEntrada)

	r.POST("/lista-espera", middleware.AuthMiddleware(), controllers.AnotarseListaEspera)

	r.GET("/mis-listas-espera", middleware.AuthMiddleware(), controllers.MisListasEspera)

	r.DELETE("/lista-espera/:id", middleware.AuthMiddleware(), controllers.SalirListaEspera)

	r.GET("/lista-espera/:conciertoId", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.VerListaEsperaConcierto)

	r.PUT("/entradas/:id/transferir", middleware.AuthMiddleware(), controllers.TransferirEntrada)

	r.GET("/reporte-ocupacion", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.ReporteOcupacion)
}
