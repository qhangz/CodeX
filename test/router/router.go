package router

import (
    "github.com/gin-gonic/gin"
    "github.com/codex/controller"
)

func InitRouter() *gin.Engine {
    router := gin.Default()
    router.GET("/users/:username", controller.GetUserByUsername)
    return router
}