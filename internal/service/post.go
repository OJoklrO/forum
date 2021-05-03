package service

import (
	"errors"
	"forum/internal/model"
)

type GetPostRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) GetPost(param *GetPostRequest) (*model.Post, error) {
	p := &model.Post{
		ID: param.ID,
	}
	err := p.Get(svc.db)
	return p, err
}

func (svc *Service) GetPostList(page, pageSize int, filter string, imageMode bool, userId string) ([]*model.Post, int, error) {
	p := model.Post{}
	pageOffset := 0
	if page > 0 {
		pageOffset = (page - 1) * pageSize
	}
	posts, err := p.List(svc.db, pageOffset, pageSize, filter, imageMode, userId)
	if err != nil {
		return nil, 0, err
	}
	count, err := p.CountAll(svc.db, filter, imageMode, userId)
	if err != nil {
		return nil, 0, err
	}

	return posts, count, nil
}

type CreatePostRequest struct {
	Title   string `json:"title" form:"title" binding:"required,min=1,max=150"`
	Content string `json:"content" form:"content" binding:"required,min=1"`
}

func (svc *Service) CreatePost(param *CreatePostRequest) (post *model.Post, err error) {
	userId := svc.ctx.Value("user_id").(string)
	p := model.Post{
		UserID: userId,
		Title:  param.Title,
	}
	post, err = p.Create(svc.db)
	if err != nil {
		return
	}
	_, err = svc.CreateComment(&CreateCommentRequest{
		PostID:  post.ID,
		Content: param.Content,
	})
	return
}

type DeletePostRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) DeletePost(param *DeletePostRequest) error {
	p, err := svc.GetPost(&GetPostRequest{ID: param.ID})
	if err != nil {
		return err
	}
	err = svc.CheckCommonPermission(p.UserID)
	if err != nil {
		return err
	}

	p = &model.Post{
		ID:    param.ID,
		IsDel: 1,
	}
	err = p.Update(svc.db)
	if err != nil {
		return err
	}

	comment := &model.Comment{PostID: p.ID, IsDel: true}
	return comment.UpdateAllByPost(svc.db)
}

func (svc *Service) SetPostPinned(id uint32) error {
	admin, err := svc.CheckAdminPermission()
	if err != nil {
		return err
	}
	if !admin {
		return errors.New("permission denied, you are not an admin")
	}

	post := &model.Post{ID: id}
	err = post.Get(svc.db)
	if err != nil {
		return err
	}

	if post.Pinned == 1 {
		post.Pinned = 2
	} else {
		post.Pinned = 1
	}

	return post.Update(svc.db)
}
