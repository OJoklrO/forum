package service

import (
	"forum/internal/model"
)

type ListCommentRequest struct {
	PostID uint32 `form:"post_id" binding:"required,gte=1"`
}

func (svc *Service) CountComments(param *ListCommentRequest) (int, error) {
	c := model.Comment{
		PostID: param.PostID,
	}
	return c.Count(svc.db)
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
	PostID  uint32 `form:"post_id" binding:"gte=0" json:"post_id"`
	Content string `form:"content" binding:"required" json:"content"`
}

func (svc *Service) CreateComment(param *CreateCommentRequest) error {
	c := model.Comment{
		UserID:  svc.ctx.Value("user_id").(string),
		PostID:  param.PostID,
		Content: param.Content,
	}
	return c.Create(svc.db)
}

type EditCommentRequest struct {
	PostID  uint32 `form:"post_id" binding:"gte=0" json:"post_id"`
	ID      uint32 `form:"id" binding:"gte=0" json:"id"`
	Content string `form:"content" binding:"required" json:"content"`
}

func (svc *Service) EditComment(param *EditCommentRequest) error {
	c := model.Comment{
		PostID:  param.PostID,
		ID:      param.ID,
		Content: param.Content,
	}
	return c.Update(svc.db)
}

type LocateCommentRequest struct {
	ID     uint32 `form:"id" binding:"required,gte=1"`
	PostID uint32 `form:"post_id" binding:"required,gte=1"`
}

func (svc *Service) DeleteComment(param *LocateCommentRequest) error {
	comment, err := svc.GetComment(param)
	if err != nil {
		return err
	}
	err = svc.CheckCommonPermission(comment.UserID)
	if err != nil {
		return err
	}
	return comment.Delete(svc.db)
}

func (svc *Service) GetComment(param *LocateCommentRequest) (*model.Comment, error) {
	c := model.Comment{
		ID:     param.ID,
		PostID: param.PostID,
	}
	return c.Get(svc.db)
}
