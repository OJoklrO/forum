package service

import (
	"github.com/OJoklrO/forum/global"
	"github.com/OJoklrO/forum/internal/model"
	"github.com/OJoklrO/forum/pkg/app"
)

type CountPostsRequest struct {
	Filter string `form:"filter"`
}

type GetPostRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type GetPostListRequest struct {
	Filter string `form:"filter"`
}

type CreatePostRequest struct {
	Title     string `form:"title" binding:"required,min=3,max=100"`
	DescImg   string `form:"desc_img" binding:"required,max=255"`
	Content   string `form:"content" binding:"required"`
	CreatedBy string `form:"created_by"`
}

type DeletePostRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountPosts(param *CountPostsRequest) (int, error) {
	return svc.dao.CountPosts(param.Filter)
}

func (svc *Service) GetPost(param *GetPostRequest) (*model.Post, error) {
	return svc.dao.GetPost(param.ID)
}

func (svc *Service) GetPostList(param *GetPostListRequest, pager *app.Pager) ([]*model.Post, error) {
	return svc.dao.GetPostList(pager.Page, pager.PageSize, param.Filter)
}

func (svc *Service) CreatePost(param *CreatePostRequest) *model.Post {
	desc := ""
	// 假设这里去掉了图片url
	//contentTemp := param.Content

	if len(param.Content) > global.AppSetting.DefaultDescLen {
		desc = param.Content[:global.AppSetting.DefaultDescLen]
	}
	return svc.dao.CreatePost(param.Title, desc, param.DescImg, param.Content, param.CreatedBy)
}

func (svc *Service) DeletePost(param *DeletePostRequest) error {
	return svc.dao.DeletePost(param.ID)
}