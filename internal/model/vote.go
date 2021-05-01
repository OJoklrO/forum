package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Vote struct {
	UserID    string `json:"user_id"`
	PostID    uint32 `json:"post_id"`
	CommentID uint32 `json:"comment_id"`
	Vote      bool   `json:"vote"`
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
