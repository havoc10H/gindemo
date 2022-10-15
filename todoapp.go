package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	router := gin.Default()

	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})

	router.POST("/todos", func(c *gin.Context) {
		var todo Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON data"})
			return
		}

		db.Create(&todo)

		c.JSON(200, todo)
	})

	router.GET("/todos/:id", func(c *gin.Context) {
		var todo Todo
		todoID := c.Param("id")

		result := db.First(&todo, todoID)
		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}

		c.JSON(200, todo)
	})

	router.PUT("/todos/:id", func(c *gin.Context) {
		var todo Todo
		todoID := c.Param("id")

		result := db.First(&todo, todoID)
		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}

		var updatedTodo Todo
		if err := c.ShouldBindJSON(&updatedTodo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON data"})
			return
		}

		todo.Title = updatedTodo.Title
		todo.Description = updatedTodo.Description
		db.Save(&todo)

		c.JSON(200, todo)
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		var todo Todo
		todoID := c.Param("id")

		result := db.First(&todo, todoID)
		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}

		db.Delete(&todo)

		c.JSON(200, gin.H{"message": fmt.Sprintf("Todo with ID %s deleted", todoID)})
	})

	router.Run(":8080")
}
