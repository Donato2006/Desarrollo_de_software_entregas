package main

import (
	"backend/dao"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dao.Connect()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
