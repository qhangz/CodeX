package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/codex/model"
	"github.com/codex/service"
)

// @Title			PublishDiscuss
// @Description		publish discuss
// @FormData		author  	formData	string		true	"作者名"
// @FormData		title   	formData	string		true	"标题"
// @FormData		summary		formData	string		true	"摘要"
// @FormData		content		formData	string		true	"内容"
// @FormData		category	formData	string		true	"标签"
// @Success			200			object		controllers.Response	"publish success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			discuss
// @Router			/api/discuss/publish [post]
func PublishDiscuss(c *gin.Context) {
	// get discuss info from request
	newDiscuss := model.Discuss{
		Author:   c.PostForm("author"),
		Title:    c.PostForm("title"),
		Summary:  c.PostForm("summary"),
		Content:  c.PostForm("content"),
		Category: c.PostForm("category"),
	}

	err := service.PublishDiscuss(newDiscuss)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "publish success",
	})
}

// @Title			GetDiscussInfo
// @Description		get discuss info by discuss id from request
// @Query			discussID  	query    	string		true 	    "id"
// @Success			200			object		controllers.Response	"success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			discuss
// @Router			/api/discuss/info [get]
func GetDiscussInfo(c *gin.Context) {
	// 将discussID转换为无符号整数类型
	discussID, strconEerr := strconv.ParseUint(c.Query("id"), 10, 64)
	if strconEerr != nil {
		// 处理转换错误
		c.JSON(http.StatusOK, gin.H{
			"code":      "500",
			"error":     strconEerr.Error(),
		})
		return
	}
	thisDiscussID := uint(discussID)
	commentList, err := service.GetDiscussInfo(thisDiscussID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": commentList,
	})
}

// @Title			Get discuss list
// @Description		get discuss list by discuss category from request
// @Query			category  	query   	string		true	    "discuss category"
// @Success			200			object		controllers.Response	"success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			discuss
// @Router			/api/discuss/list [get]
func GetDiscussList(c *gin.Context) {
	category := c.Query("category")
	discussList, err := service.GetDiscussList(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": discussList,
	})
}

// @Title			Get the top topic(title) of discuss list
// @Description		get discuss list by discuss category from request
// @Query			category  	query   	string		true	    "form pre to end"
// @Success			200			object		controllers.Response	"success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			discuss
// @Router			/api/discuss/toplist [get]
func GetDiscussTop(c *gin.Context) {
	preStr := c.Query("pre")
	endStr := c.Query("end")
	pre, preErr := strconv.Atoi(preStr)
	if preErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "400",
			"error": preErr.Error(),
		})
		return
	}
	end, endErr := strconv.Atoi(endStr)
	if endErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "400",
			"error": endErr.Error(),
		})
		return
	}
	topDiscussList, err := service.GetDiscussTop(pre, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": topDiscussList,
	})
}

// @Title			Add discuss view number
// @Description		add discuss view number by discuss id from request
// @Query			id  	    query   	string		true	    "id"
// @Success			200			object		controllers.Response	"success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			discuss
// @Router			/api/discuss/info/view [put]
func AddDiscussView(c *gin.Context) {
	idStr := c.Query("id")
	id, idErr := strconv.Atoi(idStr)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "400",
			"error": idErr.Error(),
		})
		return
	}
	err := service.AddDiscussView(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
	})
}

// @Title			Add discuss like number
// @Description		add discuss like number by discuss id from request
// @Query			id  	    query   	string		true	    "id"
// @Success			200			object		controllers.Response	"success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			discuss
// @Router			/api/discuss/info/like [put]
func AddDiscussLike(c *gin.Context) {
	idStr := c.Query("id")
	id, idErr := strconv.Atoi(idStr)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "400",
			"error": idErr.Error(),
		})
		return
	}
	err := service.AddDiscussLike(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
	})
}