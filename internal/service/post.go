package service

import (
	"forum/internal/model"
)

func (svc *Service) CountPosts() (int, error) {
	p := model.Post{}
	return p.CountAll(svc.db)
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
	err = svc.CreateComment(&CreateCommentRequest{
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
	return p.Update(svc.db)
}
