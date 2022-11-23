package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/register", register)
	r.POST("/login", login)
	UserGroup := r.Group("/user")
	{
		UserGroup.POST("/resetPassword", findPassword, remakePassword)
		UserGroup.POST("/reset_mibao", findPassword, remakeQuestion)
		UserGroup.GET("/readComment", readComments)
		UserGroup.POST("/writeComment", comment)
	}
	r.Run()

}
