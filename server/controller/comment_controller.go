package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/codex/model"
	"github.com/codex/service"
)

// @Title			PublishComment
// @Description		publish comment
// @Param			author  	formData	string		true	"作者名"
// @Param			discussID	formData	int 		true	"discussID"
// @Param			content		formData	string		true	"内容"
// @Success			200			object		controllers.Response	"publish success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			discuss
// @Router			/api/discuss/poublish [post]
func PublishComment(c *gin.Context) {
	// get discuss info from request
	// newComment := model.Comment{
	// 	Author:    c.PostForm("author"),
	// 	// DiscussID: c.PostForm("discussID"),
	// 	Content:   c.PostForm("content"),
	// }
	newComment := model.Comment{
		Author:    c.PostForm("author"),
		DiscussID: 0,
		Content:   c.PostForm("content"),
		// DiscussID: c.PostForm("discussID")
	}
	// 将评论ID转换为无符号整数类型
	discussID, err := strconv.ParseUint(c.PostForm("discussid"), 10, 64)
	if err != nil {
		// 处理转换错误
		c.JSON(http.StatusOK, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	newComment.DiscussID = uint(discussID)

	publishErr := service.PublishComment(newComment)
	if publishErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "400",
			"error": publishErr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "publish success",
	})
}

// @Title			Add comment view number
// @Description		add comment view number by comment id from request
// @Query			id  	    query   	string		true	    "id"
// @Success			200			object		controllers.Response	"success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			comment
// @Router			/api/comment/info/view [put]
func AddCommentView(c *gin.Context) {
	idStr := c.Query("id")
	id, idErr := strconv.Atoi(idStr)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "400",
			"error": idErr.Error(),
		})
		return
	}
	err := service.AddCommentView(uint(id))
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

// @Title			Add comment like number
// @Description		add comment like number by comment id from request
// @Query			id  	    query   	string		true	    "id"
// @Success			200			object		controllers.Response	"success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			comment
// @Router			/api/comment/info/like [put]
func AddCommentLike(c *gin.Context) {
	idStr := c.Query("id")
	id, idErr := strconv.Atoi(idStr)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  "400",
			"error": idErr.Error(),
		})
		return
	}
	err := service.AddCommentLike(uint(id))
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