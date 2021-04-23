package service

import (
	"context"
	"github.com/OJoklrO/forum/global"
	"github.com/OJoklrO/forum/internal/dao"
	"github.com/jinzhu/gorm"
)

type Service struct {
	ctx context.Context
	// todo: fuck dao
	dao *dao.Dao
	db  *gorm.DB
}

func New(ctx context.Context) Service {
	svc := Service{ctx, dao.New(global.DBEngine), global.DBEngine}
	return svc
}
