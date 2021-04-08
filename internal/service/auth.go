package service

import "errors"

type AuthRequest struct {
	Uname    string `form:"uname" binding:"required"`
	Upassword string `form:"upassword" binding:"required"`
}

type CreateAuthRequest struct {
	Uname    string `form:"uname" binding:"required"`
	Upassword string `form:"upassword" binding:"required"`
}

type DeleteAuthRequest struct {
	ID uint32 `form:"id" binding:"required"`
}

type AuthExistRequest struct {
	Uname string
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(
		param.Uname,
		param.Upassword,
	)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist.")
}

func (svc *Service) CreateAuth(param *CreateAuthRequest) error {
	return svc.dao.CreateAuth(param.Uname, param.Upassword)
}

func (svc *Service) DeleteAuth(param *DeleteAuthRequest) error {
	return svc.dao.DeleteAuth(param.ID)
}

func (svc *Service) AuthExist(param *AuthExistRequest) bool {
	return svc.dao.AuthExist(param.Uname)
}