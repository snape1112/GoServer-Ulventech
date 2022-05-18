package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/ithunter101/GoServer-Ulventech/ServerTest/src/controllers"
)

func main() {
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./src/pages", true)))

	// Routes
	router.POST("/parse_book", controllers.Parsebook)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// Run the server
	router.Run()
}
