package main

import (
	"backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "Backend funcionando y conectado a MySQL",
		})
	})

	r.Run(":8080")
}
