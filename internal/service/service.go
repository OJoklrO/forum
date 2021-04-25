package service

import (
	"forum/global"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Service struct {
	ctx *gin.Context
	db  *gorm.DB
}

func New(ctx *gin.Context) Service {
	svc := Service{ctx, global.DBEngine}
	return svc
}
