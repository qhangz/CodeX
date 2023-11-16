package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"

	"codex.com/codex/controllers"
	"codex.com/codex/middlewares"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.Use(middlewares.Cors())

	r.GET("/", controllers.Default())
	r.GET("/json", controllers.DefaultJson())

	user:=r.Group("/user")
	{
		user.GET("/:id", controllers.UserId())
	}

	article:=r.Group("/article")
	{
		article.GET("/list", controllers.GetArticleList())
	}

	r.Run(":8080")
}
