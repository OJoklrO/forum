package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"forum/internal/model"
	"forum/pkg/app"
)

type LoginRequest struct {
	ID       string `form:"id" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (svc *Service) LoginAccount(param *LoginRequest) (token string, err error) {
	param.Password = fmt.Sprintf("%x", md5.Sum([]byte(param.Password)))
	account := model.Account{ID: param.ID, Password: param.Password}
	err = account.CheckPassword(svc.db)
	if err != nil {
		return "", err
	}

	token, err = app.GenerateJWTToken(param.ID, param.Password)
	if err != nil {
		return "", err
	}

	return token, nil
}

type RegisterRequest struct {
	ID         string `form:"id" binding:"required"`
	Password   string `form:"password" binding:"required"`
	Name       string `form:"name" binding:"required"`
	InviteCode string `form:"invite_code" binding:"required" json:"invite_code"`
}

func (svc *Service) RegisterAccount(param *RegisterRequest) (token string, err error) {
	account := model.Account{ID: param.ID}
	exist, err := account.Exist(svc.db)
	if err != nil {
		return "", err
	}
	if exist {
		return "", errors.New("the id is not available")
	}

	param.Password = fmt.Sprintf("%x", md5.Sum([]byte(param.Password)))

	auth := model.Account{
		ID:       param.ID,
		Password: param.Password,
		Name:     param.Name,
	}

	err = auth.Create(svc.db)
	if err != nil {
		return "", err
	}

	token, err = app.GenerateJWTToken(param.ID, param.Password)
	if err != nil {
		return "", err
	}

	return token, nil
}

type DeleteAccountRequest struct {
	ID string `form:"id" binding:"required"`
}

func (svc *Service) DeleteAccount(param *DeleteAccountRequest) error {
	err := svc.CheckCommonPermission(param.ID)
	if err != nil {
		return err
	}
	auth := model.Account{
		ID: param.ID,
	}
	exist, err := auth.Exist(svc.db)
	if err != nil || !exist {
		return errors.New("account not exist")
	}
	return auth.Delete(svc.db)
}

func (svc *Service) IsAdmin(id string) (bool, error) {
	admin := model.Admin{ID: id}
	return admin.Exist(svc.db)
}

func (svc *Service) CheckAdminPermission() (bool, error) {
	id := svc.ctx.Value("user_id").(string)
	admin := model.Admin{ID: id}
	return admin.Exist(svc.db)
}

func (svc *Service) CheckCommonPermission(ownerId string) error {
	if ownerId != svc.ctx.Value("user_id").(string) {
		isAdmin, err := svc.CheckAdminPermission()
		if err != nil {
			return err
		}
		if !isAdmin {
			return errors.New("permission denied")
		}
	}
	return nil
}

type ResetPasswordRequest struct {
	OldPassword string `form:"old_password" binding:"required"`
	NewPassword string `form:"new_password" binding:"required"`
}

func (svc *Service) ResetPassword(param *ResetPasswordRequest) (err error) {
	param.OldPassword = fmt.Sprintf("%x", md5.Sum([]byte(param.OldPassword)))
	param.NewPassword = fmt.Sprintf("%x", md5.Sum([]byte(param.NewPassword)))
	account := model.Account{ID: svc.ctx.Value("user_id").(string), Password: param.OldPassword}
	err = account.CheckPassword(svc.db)
	if err != nil {
		return err
	}

	account.Password = param.NewPassword
	return account.Update(svc.db)
}
