package service

import (
	"forum/internal/model"
)

func (svc *Service) CountPosts() (int, error) {
	p := model.Post{}
	return p.Count(svc.db)
}

type GetPostRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) GetPost(param *GetPostRequest) (*model.Post, error) {
	p := model.Post{
		ID: param.ID,
	}
	return p.Get(svc.db)
}

func (svc *Service) GetPostList(page, pageSize int) ([]*model.Post, error) {
	p := model.Post{}
	pageOffset := 0
	if page > 0 {
		pageOffset = (page - 1) * pageSize
	}
	return p.List(svc.db, pageOffset, pageSize)
}

type CreatePostRequest struct {
	Title  string `json:"title" form:"title" binding:"required,min=1,max=100"`
	UserID uint32 `json:"user_id" form:"user_id" binding:"required,gte=1"`
}

func (svc *Service) CreatePost(param *CreatePostRequest) *model.Post {
	p := model.Post{
		UserID: param.UserID,
		Title:  param.Title,
	}
	return p.Create(svc.db)
}

type DeletePostRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) DeletePost(param *DeletePostRequest) error {
	p := model.Post{
		ID:    param.ID,
		IsDel: 1,
	}
	return p.Update(svc.db, p)
}
