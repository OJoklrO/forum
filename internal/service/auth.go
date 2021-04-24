package service

import (
	"errors"
	"forum/internal/model"
)

type LoginRequest struct {
	ID       string `form:"id" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (svc *Service) LoginAccount(param *LoginRequest) error {
	account := model.Account{ID: param.ID, Password: param.Password}
	return account.Check(svc.db)
}

type RegisterRequest struct {
	ID       string `form:"id" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (svc *Service) RegisterAccount(param *RegisterRequest) error {
	account := model.Account{ID: param.ID}
	exist, err := account.Exist(svc.db)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("the id is not available")
	}

	auth := model.Account{
		ID:       param.ID,
		Password: param.Password,
	}
	return auth.Create(svc.db)
}

type DeleteAccountRequest struct {
	ID string `form:"id" binding:"required"`
}

func (svc *Service) DeleteAccount(param *DeleteAccountRequest) error {
	auth := model.Account{
		ID: param.ID,
	}
	return auth.Delete(svc.db)
}
