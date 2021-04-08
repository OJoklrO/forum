package dao

import (
	"github.com/OJoklrO/forum/internal/model"
	"github.com/OJoklrO/forum/pkg/app"
)

func (d *Dao) CountComments(postID uint32) (int, error) {
	c := model.Comment{
		PostID: postID,
	}
	return c.Count(d.engine)
}

func (d *Dao) ListComment(postID uint32, page, pageSize int) ([]*model.Comment, error) {
	c := model.Comment{
		PostID: postID,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return c.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateComment(content string, postID uint32, createdBy string) error {
	c := model.Comment{
		Model: &model.Model{
			CreatedBy: createdBy,
		},
		Content: content,
		PostID: postID,
	}
	return c.Create(d.engine)
}

func (d *Dao) DeleteComment(id uint32) error {
	c := model.Comment{
		Model: &model.Model{ID: id},
	}
	return c.Delete(d.engine)
}