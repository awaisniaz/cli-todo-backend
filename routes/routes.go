package routes

import (
	"github.com/cli-todo/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.PATCH("/tasks/:id/complete", controllers.CompleteTask)
	router.DELETE("/tasks/:id", controllers.CreateTask)
	return router
}
