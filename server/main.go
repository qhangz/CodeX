package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"

	"github.com/codex/controllers"
	"github.com/codex/middlewares"
)

func main() {
	fmt.Println("Hello World")

	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(middlewares.Cors())

	router.GET("/", controllers.Default())
	router.GET("/json", controllers.DefaultJson())

	user := router.Group("/user")
	{
		user.GET("/:id", controllers.UserId())
	}

	article := router.Group("/article")
	{
		article.GET("/hot", controllers.GetArticleHot())
		article.GET("/follow", controllers.GetArticleFollow())
		article.GET("/frontend", controllers.GetArticleFrontend())
		article.GET("/backend", controllers.GetArticleBackend())
	}

	router.Run(":8080")
}
