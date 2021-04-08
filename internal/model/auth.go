package model

import (
	"github.com/jinzhu/gorm"
)

type Auth struct {
	*Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (a Auth) TableName() string {
	return "user"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("name = ? AND password = ? AND is_del = ?", a.Name, a.Password, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}

func (a Auth) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Auth) Delete(db *gorm.DB) error {
	return db.Model(&Auth{}).Where("name = ?", a.Name).Delete(&a).Error
}

func (a Auth) Exist(db *gorm.DB) bool {
	var count int
	err := db.Model(&Auth{}).Where("name = ?", a.Name).Count(&count).Error
	if err != nil {
		return false
	}
	return count < 1
}
