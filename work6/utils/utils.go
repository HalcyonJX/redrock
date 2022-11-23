package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
	})
}

func RespFail(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  500,
		"message": message,
	})
}
func RespComment(c *gin.Context, yourname, content, myname string) {
	c.JSON(http.StatusOK, gin.H{
		"留言对象": yourname,
		"留言":   content,
		"留言人":  myname,
	})
}
