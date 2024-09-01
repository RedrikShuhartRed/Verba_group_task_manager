package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/RedrikShuhartRed/TaskManager/api/handlers"
	"github.com/RedrikShuhartRed/TaskManager/api/service"
)

// RegisterRoutes registers the routes for handling task-related HTTP requests.
var RegisterRoutes = func(router *gin.Engine, service *service.TaskService) {
	handler := handlers.NewHandler(service)
	router.POST("/tasks", handler.AddNewTask)
	router.GET("/tasks", handler.GetAllTasks)
	router.GET("/tasks/:id", handler.GetTaskByID)
	router.PUT("/tasks/:id", handler.UpdateTaskByID)
	router.DELETE("/tasks/:id", handler.DeleteTask)
}
