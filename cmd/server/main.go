package main

import (
	"CLY_TodoApp_BE/internal/handlers"
	"CLY_TodoApp_BE/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupDefaultRoutes(router)

	router.Run(":8000")
}
