package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Message struct {
	From      string `json:"from"`
	To        string `json:"to"`
	PostID    uint32 `json:"post_id"`
	CommentID uint32 `json:"comment_id"`
	Read      bool   `json:"read"`
}

func (m *Message) TableName() string {
	return "message"
}

func (m *Message) Create(db *gorm.DB) error {
	// todo: at others only once
	return db.Create(m).Error
}

func (m *Message) Update(db *gorm.DB) error {
	target := *m
	target.Read = false
	return db.Model(m).Where(target).Update("read", m.Read).Error
}

func (m *Message) CountList(db *gorm.DB, filter string) (count int, err error) {
	err = db.Model(m).Where("`to` = ? AND `from` LIKE ?", m.To, "%"+filter+"%").Count(&count).Error
	if err != nil {
		fmt.Println("no ", err.Error())
	}
	return
}

func (m *Message) GetList(db *gorm.DB, pageOffset, pageSize int, filter string) (results []*Message, err error) {
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	err = db.Model(m).Where("`to` = ? AND `from` LIKE '%"+filter+"%' ", m.To).
		Order("`read`").
		Find(&results).Error
	return
}

func (m *Message) CountUnread(db *gorm.DB) (int, error) {
	res := 0
	err := db.Model(m).Where(map[string]interface{}{
		"to":   m.To,
		"read": false,
	}).Count(&res).Error
	return res, err
}
