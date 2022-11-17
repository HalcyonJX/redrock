package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock/work5/dao"
	"redrock/work5/model"
)

func register(c *gin.Context) {
	//表单验证
	if err := c.ShouldBind(&model.User{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "验证失败",
		})
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")

	flag := dao.SelectUser(username)
	if flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "此用户已存在",
		})
		return
	}
	dao.AddUser(username, password)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "注册成功",
	})
}
func login(c *gin.Context) {
	//表单验证
	if err := c.ShouldBind(&model.User{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "验证失败",
		})
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")

	flag := dao.SelectUser(username)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "此用户不存在",
		})
		return
	}
	rightPassword := dao.SelectUserPassword(username)
	if password != rightPassword {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "密码错误",
		})
		return
	}
	c.SetCookie("user_cookie", "test", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "登录成功",
	})
	return
}
