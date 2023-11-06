package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"codex.com/codex/middlewares"
	"net/http"
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
	// 动态路由(获取路由param参数)
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 获取query参数
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "person")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	// 获取post参数
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "123456") // 设置默认值
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// query和post混合参数
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "123456")

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	// Map参数（字典参数）
	r.POST("/postMap", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// 重定向
	r.GET("/redirect",func (c *gin.Context){
		c.Redirect(http.StatusFound,"http://www.bing.com")	// 302临时重定向
		// c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")	// 301永久重定向
	})

	r.Run(":8080")
}
