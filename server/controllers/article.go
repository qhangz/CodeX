package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
)

func GetArticleList() gin.HandlerFunc{
	return func(ctx *gin.Context){
		data,err:=ioutil.ReadFile("./data/article_list.json")
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to read file"})
			return 
		}
		ctx.Data(http.StatusOK,"application/json",data)
	}
}