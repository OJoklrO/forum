package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	ID     uint32 `gorm:"primary_key" json:"id"`
	Title  string `json:"title"`
	UserID uint32 `json:"user_id"`
	IsDel  uint8  `json:"is_del"`
}

func (p Post) TableName() string {
	return "post"
}

func (p Post) Count(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&p).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p Post) Get(db *gorm.DB) (*Post, error) {
	var post Post
	err := db.Model(post).Where("id = ? AND is_del = ?", p.ID, 0).Find(&post).Error
	return &post, err
}

func (p Post) List(db *gorm.DB, pageOffset, pageSize int) ([]*Post, error) {
	var posts []*Post
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if err = db.Model(Post{}).Where("is_del = ?", 0).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (p Post) Create(db *gorm.DB) *Post {
	return db.Create(&p).Value.(*Post)
}

func (p Post) Update(db *gorm.DB, v interface{}) error {
	return db.Model(&Post{}).Where("id = ?", p.ID).Updates(v).Error
}
