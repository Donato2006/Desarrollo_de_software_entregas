package main

import (
	database "backend/dao"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "Backend funcionando",
		})
	})

	routes.SetupRoutes(r)

	r.Run(":8080")
}
