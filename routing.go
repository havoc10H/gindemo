package main

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

type UserController struct{}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")

	c.JSON(200, gin.H{"id": userID, "name": "John Doe", "email": "john@example.com"})
}

func main() {
	router := gin.Default()

	userController := &UserController{}

	router.GET("/users/:id", userController.GetUserInfo)

	// public := router.Group("/public")
	// {
	// 	public.GET("/info", func(c *gin.Context) {
	// 		c.String(200, "Plublic information")
	// 	})
	// 	public.GET("/products", func(c *gin.Context) {
	// 		c.String(200, "Public product list")
	// 	})
	// }

	// private := router.Group("/private")
	// private.Use(AuthMiddleware())
	// {
	// 	private.GET("/data", func(c *gin.Context) {
	// 		c.String(200, "Private data accessible after authentication")
	// 	})
	// 	private.POST("/create", func(c *gin.Context) {
	// 		c.String(200, "Create a new resource")
	// 	})
	// }

	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello, World!")
	// })

	// router.GET("/users/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	c.String(200, "User ID: "+id)
	// })

	// router.GET("/search", func(c *gin.Context) {
	// 	query := c.DefaultQuery("q", "default-value")
	// 	c.String(200, "Search query: "+query)
	// })

	router.Run(":8080")
}
