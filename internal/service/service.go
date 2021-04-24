package service

import (
	"context"
	"forum/global"
	"github.com/jinzhu/gorm"
)

type Service struct {
	ctx context.Context
	db  *gorm.DB
}

func New(ctx context.Context) Service {
	svc := Service{ctx, global.DBEngine}
	return svc
}
