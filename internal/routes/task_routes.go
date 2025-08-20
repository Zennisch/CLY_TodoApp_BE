package routes

import (
	"CLY_TodoApp_BE/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(v1 *gin.RouterGroup, taskHandler *handlers.TaskHandler) {
	tasks := v1.Group("/tasks")
	{
		tasks.POST("", taskHandler.CreateTask)
		tasks.GET("", taskHandler.GetTasks)
	}
}
