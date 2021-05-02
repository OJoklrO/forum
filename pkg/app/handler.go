package app

import (
	"forum/global"
	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, code int, msg string) {
	global.Logger.Error(msg)
	c.JSON(code, gin.H{
		"msg": msg,
	})
	return
}
