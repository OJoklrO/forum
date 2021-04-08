package model

type Comment struct {
	*Model
	Content string `json:"content"`
	Agree int `json:"agree"`
	Disagree int `json:"disagree"`
}

func (c Comment) TableName() string {
	return "comment"
}