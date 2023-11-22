package router

import (
	"os/user"

	api "github.com/codex/api/controllers"
	"github.com/codex/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{

	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(middlewares.Cors())

	router.GET("/", api.Default())
	router.GET("/json", api.DefaultJson())

	user := router.Group("/user")
	{
		user.GET("/:id", api.UserId())
	}

	article := router.Group("/article")
	{
		article.GET("/hot", api.GetArticleHot())
		article.GET("/follow", api.GetArticleFollow())
		article.GET("/frontend", api.GetArticleFrontend())
		article.GET("/backend", api.GetArticleBackend())
	}

	// apiv1:=router.Group("/api/v1")
	// {
	// 	apiv1.GET("/tags", api.GetTag)
	// 	apiv1.POST("/tags", api.AddTag)
	// 	apiv1.PUT("/tags/:id", api.EditTag)
	// 	apiv1.DELETE("/tags/:id", api.DeleteTag)
	// }

	return router
}