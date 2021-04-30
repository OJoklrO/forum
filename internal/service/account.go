package service

import "forum/internal/model"

type AccountInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Gender      uint32 `json:"gender"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Level       uint32 `json:"level"`
}

func (svc *Service) GetUserInfo(id string) (*AccountInfo, error) {
	res := model.Account{
		ID: id,
	}
	err := res.Get(svc.db)
	if err != nil {
		return nil, err
	}

	return &AccountInfo{
		ID:          res.ID,
		Name:        res.Name,
		Gender:      res.Gender,
		Avatar:      res.Avatar,
		Description: res.Description,
		Level:       res.Level,
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
