package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"codex.com/codex/middlewares"
)

func main() {
	fmt.Println("Hello World")
	
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/", func(c *gin.Context) {
		// c.String(200, "Hello myGin!")
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.Run(":8080")
}
