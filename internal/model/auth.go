package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Account struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	IsDel    bool   `json:"is_del"`
}

func (a Account) TableName() string {
	return "account"
}

func (a Account) Check(db *gorm.DB) error {
	var account Account
	db = db.Where("id = ? AND password = ? AND is_del = ?",
		a.ID, a.Password, 0)
	return db.First(&account).Error
}

func (a Account) Exist(db *gorm.DB) (bool, error) {
	var count int
	err := db.Model(&Account{}).Where("id = ?", a.ID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a Account) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Account) Delete(db *gorm.DB) error {
	exist, err := a.Exist(db)
	if !exist || err != nil {
		return errors.New("the account does not exists")
	}
	a.IsDel = true
	return db.Model(&Account{}).Where("id = ?", a.ID).Update(&a).Error
}
