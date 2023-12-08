package router

import (
	// "os/user"

	"github.com/codex/controller"
	"github.com/codex/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// gin.SetMode(gin.DebugMode)
	// router := gin.New()
	router := gin.Default()
	router.Use(middlewares.Cors())

	router.GET("/users/:username", controller.GetUserByUsername)

	article := router.Group("/api/article")
	{
		article.GET("/hot", controller.GetArticleHot())
		article.GET("/follow", controller.GetArticleFollow())
		article.GET("/frontend", controller.GetArticleFrontend())
		article.GET("/backend", controller.GetArticleBackend())
	}

	user := router.Group("/api/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}

	return router
}
