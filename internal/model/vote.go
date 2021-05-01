package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Vote struct {
	UserID    string `json:"user_id"`
	PostID    uint32 `json:"post_id"`
	CommentID uint32 `json:"comment_id"`
	Vote      int    `json:"vote"`
}

func (v *Vote) TableName() string {
	return "vote"
}

func (v *Vote) SetOrCreate(db *gorm.DB) error {
	result := db.Model(v).Where("user_id = ? AND post_id = ? AND comment_id = ?", v.UserID, v.PostID, v.CommentID).Update(map[string]interface{}{"vote": v.Vote})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) ||
		result.RowsAffected == 0 {
		return db.Create(v).Error
	} else {
		return result.Error
	}
}

func (v *Vote) Get(db *gorm.DB) error {
	return db.Model(v).Where("user_id = ? AND post_id = ? AND comment_id = ?", v.UserID, v.PostID, v.CommentID).Find(v).Error
}

type sumResult struct {
	Total int
}

func (v *Vote) Sum(db *gorm.DB) (count int, err error) {
	var sum sumResult
	result := db.Model(v).Select("sum(vote) as total").Where("post_id = ? AND comment_id = ?", v.PostID, v.CommentID).Scan(&sum)
	count = sum.Total
	err = result.Error
	return
}
