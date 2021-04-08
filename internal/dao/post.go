package dao

import (
	"github.com/OJoklrO/forum/internal/model"
	"github.com/OJoklrO/forum/pkg/app"
)

func (d *Dao) CountPosts(filter string) (int, error) {
	p := model.Post{}
	return p.Count(d.engine, filter)
}

func (d *Dao) GetPost(id uint32) (*model.Post, error) {
	p := model.Post{
		Model: &model.Model{
			ID: id,
		},
	}
	return p.Get(d.engine)
}

func (d *Dao) GetPostList(page, pageSize int, filter string) ([]*model.Post, error) {
	p := model.Post{}
	pageOffset := app.GetPageOffset(page, pageSize)
	return p.List(d.engine, pageOffset, pageSize, filter)
}

func (d *Dao) CreatePost(title, desc, descImg, content, createdBy string) error {
	p := model.Post{
		Title: title,
		Desc: desc,
		DescImg: descImg,
		Content: content,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return p.Create(d.engine)
}

func (d *Dao) DeletePost(id uint32) error {
	p := model.Post{
		Model: &model.Model{
			ID: id,
		},
	}
	return p.Delete(d.engine)
}