package main

import (
	"CLY_TodoApp_BE/internal/handlers"
	"CLY_TodoApp_BE/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupDefaultRoutes(router)

	v1 := router.Group("/api/v1")
	taskHandler := handlers.NewTaskHandler()
	routes.SetupTaskRoutes(v1, taskHandler)

	router.Run(":8000")
}
