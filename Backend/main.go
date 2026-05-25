package main

import (
	database "backend/dao"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa la conexión a la base de datos y corre las migraciones
	database.Connect()
	// Crea una instancia por defecto del motor de Gin (incluye middlewares de Logger y Recovery)
	r := gin.Default()
	// Define un endpoint rápido en la raíz "/" para chequear la salud del servidor (Health Check)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "Backend funcionando",
		})
	})
	// Carga y configura el resto de las rutas de nuestra API en el servidor Gin
	routes.SetupRoutes(r)

	r.Run(":8080")
}
