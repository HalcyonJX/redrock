package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.POST("register", register)
	r.POST("login", login)
	r.Run(":9090")
}
