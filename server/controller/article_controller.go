package controller

import (
	"github.com/gin-gonic/gin"
	// "github.com/codex/service"
	"net/http"
	"io/ioutil"
)

// func GetArticleHot(c *gin.Context) {
// 	data, err := service.GetArticleHot()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, data)
// }


func GetArticleList() gin.HandlerFunc{
	return func(ctx *gin.Context){
		data,err:=ioutil.ReadFile("./resource/article_list.json")
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to read file"})
			return 
		}
		ctx.Data(http.StatusOK,"application/json",data)
	}
}

// 热门文章
func GetArticleHot() gin.HandlerFunc{
	return func(ctx *gin.Context){
		data,err:=ioutil.ReadFile("./resource/article_hot.json")
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to read file"})
			return 
		}
		ctx.Data(http.StatusOK,"application/json",data)
	}
}

// 关注文章
func GetArticleFollow() gin.HandlerFunc{
	return func(ctx *gin.Context){
		data,err:=ioutil.ReadFile("./resource/article_follow.json")
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to read file"})
			return 
		}
		ctx.Data(http.StatusOK,"application/json",data)
	}
}

// 前端文章
func GetArticleFrontend() gin.HandlerFunc{
	return func(ctx *gin.Context){
		data,err:=ioutil.ReadFile("./resource/article_frontend.json")
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to read file"})
			return 
		}
		ctx.Data(http.StatusOK,"application/json",data)
	}
}

// 后端文章
func GetArticleBackend() gin.HandlerFunc{
	return func(ctx *gin.Context){
		data,err:=ioutil.ReadFile("./resource/article_backend.json")
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to read file"})
			return 
		}
		ctx.Data(http.StatusOK,"application/json",data)
	}
}