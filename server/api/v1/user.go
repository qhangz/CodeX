package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserId() gin.HandlerFunc{
	return func(ctx *gin.Context){
		id:=ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}