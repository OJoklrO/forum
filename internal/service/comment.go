package service

import (
	"github.com/OJoklrO/forum/internal/model"
)

type CountCommentsRequest struct {
	PostID uint32 `form:"post_id" binding:"required,gte=1"`
}

func (svc *Service) CountComments(param *CountCommentsRequest) (int, error) {
	c := model.Comment{
		PostID: param.PostID,
	}
	return c.Count(svc.db)
}

type ListCommentRequest struct {
	PostID uint32 `form:"post_id" binding:"required,gte=1"`
}

func (svc *Service) ListComment(param *ListCommentRequest, page, pageSize int) ([]*model.Comment, error) {
	c := model.Comment{PostID: param.PostID}
	pageOffset := 0
	if page > 0 {
		pageOffset = (page - 1) * pageSize
	}
	return c.List(svc.db, pageOffset, pageSize)
}

type CreateCommentRequest struct {
	UserID  string `form:"user_id" binding:"required" json:"user_id"`
	PostID  uint32 `form:"post_id" binding:"gte=0" json:"post_id"`
	Content string `form:"content" binding:"required" json:"content"`
}

func (svc *Service) CreateComment(param *CreateCommentRequest) error {
	c := model.Comment{
		UserID:  param.UserID,
		PostID:  param.PostID,
		Content: param.Content,
	}
	return c.Create(svc.db)
}

type DeleteCommentRequest struct {
	ID     uint32 `form:"id" binding:"required,gte=1"`
	PostID uint32 `form:"post_id" binding:"required,gte=1"`
}

func (svc *Service) DeleteComment(param *DeleteCommentRequest) error {
	c := model.Comment{
		ID:     param.ID,
		PostID: param.PostID,
	}
	return c.Delete(svc.db)
}
