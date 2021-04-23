package dao

import "github.com/jinzhu/gorm"

// todo: delete all dao
type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine}
}