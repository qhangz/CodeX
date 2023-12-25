package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
)

// @Title			get games list
// @Description		获取游戏列表
// @Success			200			object		controllers.Response	"code,data"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/api/game/list [get]
func GetGameList(c *gin.Context){
	data,err:=ioutil.ReadFile("./resource/game_list.json")
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "400",
			"error":"Failed to read file",
		})
		return 
	}
	c.Data(http.StatusOK,"application/json",data)
}

