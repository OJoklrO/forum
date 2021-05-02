package model

import "github.com/jinzhu/gorm"

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
	return db.Create(m).Error
}

func (m *Message) Update(db *gorm.DB) error {
	target := *m
	target.Read = false
	return db.Model(m).Where(target).Update("read", m.Read).Error
}

func (m *Message) GetUnreadList(db *gorm.DB) (results []*Message, err error) {
	err = db.Where(map[string]interface{}{
		"to":   m.To,
		"read": false,
	}).Find(&results).Error
	return
}
