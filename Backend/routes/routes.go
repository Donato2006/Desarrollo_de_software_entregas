package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura las rutas REST de la API y las vincula con el motor de Gin
func SetupRoutes(r *gin.Engine) {
	// GET /conciertos -> Devuelve la lista de todos los conciertos
	r.GET("/conciertos", controllers.ObtenerConciertos)
	// POST /conciertos -> Crea un nuevo concierto a partir del JSON enviado
	r.POST("/conciertos", controllers.CrearConcierto)
	// GET /conciertos/:id -> Devuelve la información de un concierto en base a su ID en la URL
	r.GET("/conciertos/:id", controllers.ObtenerConciertoPorID)
	// PUT /conciertos/:id -> Modifica los datos de un concierto específico mediante su ID
	r.PUT("/conciertos/:id", controllers.ActualizarConcierto)
	// DELETE /conciertos/:id -> Elimina de manera física o lógica un concierto por ID
	r.DELETE("/conciertos/:id", controllers.EliminarConcierto)
}
