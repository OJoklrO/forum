package model

import (
	"github.com/OJoklrO/forum/pkg/app"
	"github.com/jinzhu/gorm"
)

type Post struct {
	*Model
	Content string `json:"content"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	DescImg string `json:"desc_img"`
	Agree int `json:"agree"`
	Disagree int `json:"disagree"`
}

func (p Post) TableName() string {
	return "post"
}

type PostSwagger struct {
	List []*Post
	Pager *app.Pager
}

func (p Post) Count(db *gorm.DB, filter string) (int, error) {
	var count int
	err := db.Model(&p).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p Post) Get(db *gorm.DB) (*Post, error) {
	var post Post
	err := db.Where("id = ? AND is_del = ?", p.ID, 0).Find(&post).Error
	return &post, err
}

func (p Post) List(db *gorm.DB, pageOffset, pageSize int, filter string) ([]*Post, error) {
	var posts []*Post
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if err = db.Where("is_del = ?", 0).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (p Post) Create(db *gorm.DB) error {
	return db.Create(&p).Error
}

func (p Post) Update(db *gorm.DB, v interface{}) error {
	return db.Model(&Post{}).Where("id = ? AND is_del = ?", p.ID, 0).Updates(v).Error
}

func (p Post) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", p.ID, 0).Delete(&p).Error
}