package service

import (
	"github.com/OJoklrO/forum/internal/model"
	"github.com/OJoklrO/forum/pkg/app"
)

type CountCommentsRequest struct {
	PostID uint32 `form:"post_id" binding:"required,gte=1"`
}

type ListCommentRequest struct {
	PostID uint32 `form:"post_id" binding:"required,gte=1"`
}

type CreateCommentRequest struct {
	Content string `form:"content" binding:"required"`
	PostID uint32 `form:"post_id" binding:"required"`
	CreatedBy string `form:"created_by" binding:"required"`
}

type DeleteCommentRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountComments(param *CountCommentsRequest) (int, error) {
	return svc.dao.CountComments(param.PostID)
}

func (svc *Service) ListComment(param *ListCommentRequest, pager *app.Pager) ([]*model.Comment, error) {
	return svc.dao.ListComment(param.PostID, pager.Page, pager.PageSize)
}

func (svc *Service) CreateComment(param *CreateCommentRequest) error {
	return svc.dao.CreateComment(param.Content, param.PostID, param.CreatedBy)
}

func (svc *Service) DeleteComment(param *DeleteCommentRequest) error {
	return svc.dao.DeleteComment(param.ID)
}