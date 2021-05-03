package model

import (
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
	voteValue := v.Vote
	return db.Where("user_id = ? AND post_id = ? AND comment_id = ?", v.UserID, v.PostID, v.CommentID).Assign(map[string]interface{}{"vote": voteValue}).FirstOrCreate(v).Error
}

func (v *Vote) Get(db *gorm.DB) error {
	return db.Model(v).Where("user_id = ? AND post_id = ? AND comment_id = ?", v.UserID, v.PostID, v.CommentID).Find(v).Error
}

type result struct {
	Total int
}

func (v *Vote) CommentSum(db *gorm.DB) (count int, err error) {
	var sum result
	result := db.Model(v).Select("sum(vote) as total").Where("post_id = ? AND comment_id = ?", v.PostID, v.CommentID).Scan(&sum)
	count = sum.Total
	err = result.Error
	return
}

func (v *Vote) CommentVoteCount(db *gorm.DB) (countUp, countDown int, err error) {
	var count result
	result := db.Model(v).Select("count(vote) as total").Where("post_id = ? AND comment_id = ? AND vote = 1", v.PostID, v.CommentID).Scan(&count)
	countUp = count.Total
	result = db.Model(v).Select("count(vote) as total").Where("post_id = ? AND comment_id = ? AND vote = -1", v.PostID, v.CommentID).Scan(&count)
	countDown = count.Total
	err = result.Error
	return
}

func (v *Vote) UserSum(db *gorm.DB) (count int, err error) {
	var sum result
	result := db.Model(v).Select("sum(vote) as total").Where("user_id = ?", v.UserID).Scan(&sum)
	count = sum.Total
	err = result.Error
	return
}
