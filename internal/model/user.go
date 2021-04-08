package model

type User struct {
	*Model
	UName string `json:"uname"`
	UPassword string `json:"upassword"`
}

func (u User) TableName() string {
	return "user"
}