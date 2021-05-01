package service

import (
	"forum/internal/model"
	"time"
)

func (svc *Service) CountCommentsOfPost(postId uint32) (int, error) {
	c := model.Comment{
		PostID: postId,
	}
	return c.CountByPostId(svc.db)
}

type ListCommentRequest struct {
	PostID uint32 `form:"post_id" binding:"required,gte=1"`
}

func (svc *Service) CountCommentUsers(param *ListCommentRequest) (int, error) {
	c := model.Comment{
		PostID: param.PostID,
	}
	return c.CountUsers(svc.db)
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
		Time:    time.Now().Format("2006-01-02"),
	}
	err := c.Create(svc.db)
	if err != nil {
		return err
	}

	count, err := svc.CountCommentsOfPost(param.PostID)
	if err != nil {
		return err
	}

	err = svc.UpdateUserLevel()
	if err != nil {
		return err
	}

	post := model.Post{
		ID:           param.PostID,
		LatestReply:  c.Time,
		ReplyUserID:  svc.ctx.Value("user_id").(string),
		CommentCount: uint32(count),
	}
	return post.Update(svc.db)
}

type EditCommentRequest struct {
	PostID  uint32 `form:"post_id" binding:"gte=0" json:"post_id"`
	ID      uint32 `form:"id" binding:"gte=0" json:"id"`
	Content string `form:"content" binding:"required" json:"content"`
}

func (svc *Service) EditComment(param *EditCommentRequest) error {
	c := model.Comment{
		PostID:   param.PostID,
		ID:       param.ID,
		Content:  param.Content,
		Time:     time.Now().Format("2006-01-02"),
		IsEdited: true,
	}
	err := c.Update(svc.db)
	if err != nil {
		return err
	}
	post := model.Post{
		ID:          param.PostID,
		LatestReply: c.Time,
		ReplyUserID: svc.ctx.Value("user_id").(string),
	}
	return post.Update(svc.db)
}

func (svc *Service) DeleteComment(id, postId uint32) error {
	comment, err := svc.GetComment(id, postId)
	if err != nil {
		return err
	}
	err = svc.CheckCommonPermission(comment.UserID)
	if err != nil {
		return err
	}
	return comment.Delete(svc.db)
}

func (svc *Service) GetComment(id, postId uint32) (*model.Comment, error) {
	c := model.Comment{
		ID:     id,
		PostID: postId,
	}
	return c.Get(svc.db)
}

func (svc *Service) Vote(id, postId uint32, support int) error {
	v := &model.Vote{
		UserID:    svc.ctx.Value("user_id").(string),
		PostID:    postId,
		CommentID: id,
		Vote:      support,
	}

	return v.SetOrCreate(svc.db)
}

func (svc *Service) GetVotes(id, postId uint32) (int, error) {
	v := &model.Vote{
		PostID:    postId,
		CommentID: id,
	}
	return v.CommentSum(svc.db)
}
