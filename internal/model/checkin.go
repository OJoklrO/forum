package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Checkin struct {
	ID    string `json:"id"`
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
}

func (c *Checkin) TableName() string {
	return "checkin"
}

func (c *Checkin) Set(db *gorm.DB) error {
	err := db.First(c, c).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Create(c).Error
	}
	return err
}

func (c *Checkin) GetThisMonth(db *gorm.DB) (checkIns []Checkin, err error) {
	result := db.Find(&checkIns, map[string]interface{}{
		"year":  time.Now().Year(),
		"month": time.Now().Month(),
		"id":    c.ID,
	})
	err = result.Error
	return
}
