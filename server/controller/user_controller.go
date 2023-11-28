package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/codex/model"
	"github.com/codex/service"
)

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := service.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Title			user register
// @Description		使用表单用户名、密码和邮箱注册
// @Param			username	formData	string		true	"用户名"
// @Param			password	formData	string		true	"密码"
// @Param			email		formData	string		true	"邮箱账号"
// @Success			200			object		controllers.Response	"register success"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/user/register [post]
func Register(c *gin.Context) {
	// get user info from request
	newUser := model.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Email:    c.PostForm("email"),
	}
	// c.Bind(&newUser)
	err := service.Register(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "register success"})
}

// @Title			user login
// @Description		使用表单用户名和密码登录
// @Param			username	formData	string		true	"用户名"
// @Param			password	formData	string		true	"密码"
// @Success			200			object		controllers.Response	"token"
// @Failure			401			object		controllers.Response	"各种错误"
// @Failure 		500 object controllers.Response "服务器内部错误"
// @Tags			user
// @Router			/user/login [post]
func Login(c *gin.Context) {
	// get user info from request
	user := model.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}
	// c.Bind(&user)
	token, err := service.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		// "token": token
		"code": "200",
		"msg":  "登录成功",
		"data": gin.H{"token": token},
	})
}
