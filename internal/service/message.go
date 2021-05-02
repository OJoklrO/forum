package service

import "forum/internal/model"

func (svc *Service) CreateNotifyMessage(to string, postId, commentId uint32) error {
	from := svc.ctx.Value("user_id").(string)
	message := &model.Message{
		From:      from,
		To:        to,
		PostID:    postId,
		CommentID: commentId,
		Read:      false,
	}
	return message.Create(svc.db)
}

func (svc *Service) ReadMessage(to string, postId, commentId uint32) error {
	message := &model.Message{
		To:        to,
		PostID:    postId,
		CommentID: commentId,
		Read:      true,
	}
	return message.Update(svc.db)
}

func (svc *Service) GetUnreadMessages() ([]*model.Message, error) {
	message := &model.Message{
		To: svc.ctx.Value("user_id").(string),
	}
	return message.GetUnreadList(svc.db)
}
