package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type Comment struct {
	ID       uint32 `json:"id"`
	PostID   uint32 `json:"post_id"`
	UserID   string `json:"user_id"`
	Content  string `json:"content"`
	Time     int64  `json:"time"`
	IsDel    bool   `json:"is_del"`
	IsEdited bool   `json:"is_edited"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *Comment) CountByUserId(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&Comment{}).Where("user_id = ? AND is_del = ?", c.UserID, 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *Comment) CountByUser(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&Comment{}).Where("user_id = ? AND is_del = ?", c.UserID, 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *Comment) CountByPostId(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&Comment{}).Where("post_id = ? AND is_del = ?", c.PostID, 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *Comment) CountUsers(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&Comment{}).Where("post_id = ?", c.PostID).Group("user_id").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *Comment) List(db *gorm.DB, pageOffset, pageSize int, filter string) ([]*Comment, error) {
	var comments []*Comment
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	query := " is_del = 0"
	if c.PostID != 0 {
		query += " AND post_id = " + strconv.Itoa(int(c.PostID)) + " "
	}
	if len(c.UserID) != 0 {
		query += " AND user_id = '" + c.UserID + "' "
	}
	query += " AND content LIKE '%" + filter + "%' "
	if err = db.Model(Comment{}).Where(query).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *Comment) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c *Comment) Delete(db *gorm.DB) error {
	c.IsDel = true
	return c.Update(db)
}

func (c Comment) Update(db *gorm.DB) error {
	return db.Model(c).Where("id = ? AND post_id = ?", c.ID, c.PostID).Update(&c).Error
}

func (c Comment) UpdateAllByPost(db *gorm.DB) error {
	return db.Model(c).Where("post_id = ?", c.PostID).Updates(&c).Error
}

func (c Comment) Get(db *gorm.DB) (*Comment, error) {
	var comment Comment
	err := db.Model(comment).Where("id = ? AND post_id = ? AND is_del = ?", c.ID, c.PostID, 0).Find(&comment).Error
	return &comment, err
}
