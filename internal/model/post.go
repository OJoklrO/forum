package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	ID           uint32 `gorm:"primary_key" json:"id"`
	Title        string `json:"title"`
	UserID       string `json:"user_id"`
	ReplyUserID  string `json:"reply_user_id"`
	IsDel        uint8  `json:"is_del"`
	Pinned       uint8  `json:"pinned"`
	LatestReply  int64  `json:"latest_reply"`
	CommentCount uint32 `json:"comment_count"`
}

func (p Post) TableName() string {
	return "post"
}

func (p Post) CountAll(db *gorm.DB, filter string, imageOnly bool) (int, error) {
	// todo: (warning) duplicated with List()
	var count int
	imageSQL := ""
	if imageOnly {
		imageSQL = " AND comment.content LIKE '%<img%src=%>%' AND comment.id = 1"
	}
	err := db.Model(Post{}).
		Joins("join comment on comment.post_id =post.id").
		Where("post.is_del = ? AND post.title LIKE ?"+imageSQL, 0, "%"+filter+"%").
		Or("comment.is_del = ? AND comment.content LIKE ?"+imageSQL, 0, "%"+filter+"%").
		Group("post.id").
		Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *Post) CountByUser(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(p).Where("is_del = ? AND user_id = ?", 0, p.UserID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *Post) Get(db *gorm.DB) error {
	return db.Model(p).Where("id = ? AND is_del = ?", p.ID, 0).Find(p).Error
}

func (p Post) List(db *gorm.DB, pageOffset, pageSize int, filter string, imageOnly bool) ([]*Post, error) {
	var posts []*Post
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	imageSQL := ""
	if imageOnly {
		imageSQL = " AND comment.content LIKE '%<img%src=%>%' AND comment.id = 1"
	}
	if err = db.Model(Post{}).
		Joins("join comment on comment.post_id =post.id").
		Where("post.is_del = ? AND post.title LIKE ?"+imageSQL, 0, "%"+filter+"%").
		Or("comment.is_del = ? AND comment.content LIKE ?"+imageSQL, 0, "%"+filter+"%").
		Order("pinned desc, latest_reply desc").
		Group("post.id").
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (p Post) Create(db *gorm.DB) (*Post, error) {
	db = db.Create(&p)
	if db.Error != nil {
		return nil, db.Error
	}
	return db.Value.(*Post), nil
}

func (p *Post) Update(db *gorm.DB) error {
	return db.Model(&Post{}).Where("id = ?", p.ID).Updates(p).Error
}
