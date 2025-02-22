package controllers

import (
	"net/http"

	"github.com/cli-todo/database"
	"github.com/cli-todo/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func CompleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	task.Completed = true
	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task Deleted Successfully"})

}
