package model

import "github.com/jinzhu/gorm"

type Admin struct {
	ID string `json:"id"`
}

func (a Admin) TableName() string {
	return "admin"
}

func (a Admin) Exist(db *gorm.DB) (bool, error) {
	var count int
	err := db.Model(&Admin{}).Where("id = ?", a.ID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
