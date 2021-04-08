package model

type PostComment struct {
	*Model
	PostID uint32 `json:"postID"`
	CommentID uint32 `json:"commentID"`
}

func (pc PostComment) TableName() string {
	return "post_comment"
}