package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
)

func Default() gin.HandlerFunc{
	return func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	}
}

func DefaultJson() gin.HandlerFunc{
	return func(ctx *gin.Context){
		data,err:=ioutil.ReadFile("./data/test.json")
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to read file"})
			return 
		}

		ctx.Data(http.StatusOK,"application/json",data)
	}
}