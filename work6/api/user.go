package api

import (
	"github.com/gin-gonic/gin"
	"redrock/work6/dao"
	"redrock/work6/utils"
)

func register(c *gin.Context) {
	//传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	question := c.PostForm("question")
	answer := c.PostForm("answer")
	if username == "" || password == "" || question == "" || answer == "" {
		utils.RespFail(c, "验证失败")
		return
	}
	//验证用户名是否重复
	flag := dao.SelectUser(username)
	if flag {
		utils.RespFail(c, "用户已经存在")
		return
	}
	dao.AddUser(username, password, question, answer)
	utils.RespSuccess(c, "注册成功")
}
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		utils.RespFail(c, "验证失败")
		return
	}
	flag := dao.SelectUser(username)
	if !flag {
		utils.RespFail(c, "用户不存在")
		return
	}
	correct := dao.SelectPassword(username)
	if password != correct {
		utils.RespFail(c, "密码错误")
		return
	}
	c.SetCookie("gin_demo_cookie", "test", 3600, "/", "localhost", false, true)
	utils.RespSuccess(c, "登录成功")
}

// 验证密保
func findPassword(c *gin.Context) {
	username := c.PostForm("username")
	question := dao.SelectQuestion(username)
	utils.RespSuccess(c, question)
	answer := c.PostForm("answer")
	correct := dao.SelectAnswer(username)
	if answer != correct {
		utils.RespFail(c, "答案错误")
		return
	}
}

// 重置密码
func remakePassword(c *gin.Context) {
	username := c.PostForm("username")
	//输入两次新密码
	password1 := c.PostForm("password1")
	password2 := c.PostForm("password2")
	if username == "" || password1 == "" || password2 == "" {
		utils.RespFail(c, "验证失败")
		return
	}
	if password1 != password2 {
		utils.RespSuccess(c, "两次密码不同,请重新输入")
		return
	}
	dao.UpdatePassword(username, password1)
	utils.RespSuccess(c, "重置密码成功")
}

// 重置密保
func remakeQuestion(c *gin.Context) {
	username := c.PostForm("username")
	question := c.PostForm("question")
	answer := c.PostForm("answer")
	if username == "" || question == "" || answer == "" {
		utils.RespFail(c, "验证失败")
		return
	}
	dao.UpdateQuestion(username, question)
	dao.UpdateAnswer(username, answer)
	utils.RespSuccess(c, "重置密保问题成功")
}

// 查看留言
func readComments(c *gin.Context) {
	myname := c.PostForm("myname")
	yourname, content, name := dao.Select(myname)
	utils.RespComment(c, yourname, content, name)
}

// 向别人留言
func comment(c *gin.Context) {
	myname := c.PostForm("myname")
	yourname := c.PostForm("yourname")
	content := c.PostForm("content")
	if content == "" || myname == "" || yourname == "" {
		utils.RespFail(c, "不能出现内容为空")
		return
	}
	dao.AddComments(yourname, content, myname)
	utils.RespSuccess(c, "留言成功")
}
