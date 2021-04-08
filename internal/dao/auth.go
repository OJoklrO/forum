package dao

import "github.com/OJoklrO/forum/internal/model"

func (d *Dao) GetAuth(uname, upassword string) (model.Auth, error) {
	auth := model.Auth{Name: uname, Password: upassword}
	return auth.Get(d.engine)
}

func (d *Dao) CreateAuth(uname, upassword string) error {
	auth := model.Auth{
		Name:     uname,
		Password: upassword,
	}
	return auth.Create(d.engine)
}

func (d *Dao) DeleteAuth(id uint32) error {
	auth := model.Auth{
		Model: &model.Model{ID: id},
	}
	return auth.Delete(d.engine)
}

func (d *Dao) AuthExist(uname string) bool {
	auth := model.Auth{
		Name: uname,
	}
	return auth.Exist(d.engine)
}