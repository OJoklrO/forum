package service

import "forum/internal/model"

type AccountInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Gender       uint32 `json:"gender"`
	Avatar       string `json:"avatar"`
	Description  string `json:"description"`
	Level        uint32 `json:"level"`
	CommentCount int    `json:"comment_count"`
	PostCount    int    `json:"post_count"`
}

func (svc *Service) GetUserInfo(id string) (*AccountInfo, error) {
	account := model.Account{
		ID: id,
	}
	err := account.Get(svc.db)
	if err != nil {
		return nil, err
	}

	commentCount, err := svc.CountCommentsOfUser(id)
	if err != nil {
		return nil, err
	}

	postCount, err := svc.CountPostsOfUser(id)
	if err != nil {
		return nil, err
	}

	return &AccountInfo{
		ID:           account.ID,
		Name:         account.Name,
		Gender:       account.Gender,
		Avatar:       account.Avatar,
		Description:  account.Description,
		Level:        account.Level,
		CommentCount: commentCount,
		PostCount:    postCount,
	}, nil
}

type EditAccountInfoRequest struct {
	ID          string `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Gender      uint32 `json:"gender" form:"name" binding:"numeric,min=1,max=3,required"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

func (svc *Service) EditUserInfo(param *EditAccountInfoRequest) error {
	err := svc.CheckCommonPermission(param.ID)
	if err != nil {
		return err
	}
	account := model.Account{
		ID:          param.ID,
		Name:        param.Name,
		Description: param.Description,
		Gender:      param.Gender,
		Avatar:      param.Avatar,
	}

	return account.Update(svc.db)
}

func (svc *Service) CountCommentsOfUser(userId string) (int, error) {
	c := model.Comment{
		UserID: userId,
	}
	return c.CountByUserId(svc.db)
}

func (svc *Service) CountPostsOfUser(userId string) (int, error) {
	post := model.Post{
		UserID: userId,
	}
	return post.CountByUser(svc.db)
}
