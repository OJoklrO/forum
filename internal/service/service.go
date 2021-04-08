package service

import (
	"context"
	"github.com/OJoklrO/forum/global"
	"github.com/OJoklrO/forum/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}


func New(ctx context.Context) Service {
	svc := Service{ctx, dao.New(global.DBEngine)}
	return svc
}

