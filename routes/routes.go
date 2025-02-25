package routes

import (
	"github.com/cli-todo/controllers"
	"github.com/cli-todo/middlewares"
	"github.com/gin-gonic/gin"
)
func SetupRoutes() *gin.Engine {
	router := gin.Default()
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}
	taskRoute := router.Group("/task")
	taskRoute.Use(middlewares.AuthMiddleware())
	{
		taskRoute.GET("/tasks", controllers.GetTasks)
		taskRoute.POST("/tasks", controllers.CreateTask)
		taskRoute.PATCH("/tasks/:id/complete", controllers.CompleteTask)
		taskRoute.DELETE("/tasks/:id", controllers.CreateTask)
	}

	return router
}
