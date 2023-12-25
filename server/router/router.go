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

	// router.GET("/users/:username", controller.GetUserByUsername)

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
		user.GET("/list", controller.GetUserList)
		user.GET("/info/:username", controller.GetUserInfoByUsername)
		user.POST("/update/username", controller.UpdateUsername)
		user.POST("/update/password", controller.UpdatePassword)
		user.POST("/update/email", controller.UpdateEmail)
		user.POST("/update/age", controller.UpdateAge)
		user.POST("/update/summary", controller.UpdateSummary)
		user.POST("/delete", controller.DeleteUser)
	}

	discuss := router.Group("/api/discuss")
	{
		discuss.POST("/publish", controller.PublishDiscuss)
		discuss.GET("/info", controller.GetDiscussInfo)
		discuss.GET("/list", controller.GetDiscussList)
		discuss.GET("/toplist", controller.GetDiscussTop)
	}

	comment := router.Group("/api/comment")
	{
		comment.POST("/publish", controller.PublishComment)
	}

	game := router.Group("/api/game")
	{
		game.GET("/list", controller.GetGameList)
	}

	event := router.Group("/api/event")
	{
		event.GET("/list", controller.GetEventList)
	}

	return router
}
