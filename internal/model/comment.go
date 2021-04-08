package model

import (
	"github.com/OJoklrO/forum/pkg/app"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	*Model
	Content string `json:"content"`
	PostID uint32 `json:"post_id"`
	Agree int `json:"agree"`
	Disagree int `json:"disagree"`
}

type CommentSwagger struct {
	List []*Comment
	Pager *app.Pager
}

func (c Comment) TableName() string {
	return "comment"
}

func (c Comment) Count(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&Comment{}).Where("post_id = ? AND is_del = ?", c.PostID, 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c Comment) List(db *gorm.DB, pageOffset, pageSize int) ([]*Comment, error) {
	var comments []*Comment
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if err = db.Where("post_id = ? AND is_del = ?",c.PostID, 0).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (c Comment) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

//func (c Comment) Update(db *gorm.DB, v interface{}) error {
//
//}

func (c Comment) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", c.ID, 0).Delete(&c).Error
}