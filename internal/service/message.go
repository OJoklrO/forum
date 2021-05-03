package service

import "forum/internal/model"

func (svc *Service) CreateNotifyMessage(to string, postId, commentId uint32) error {
	from := svc.ctx.Value("user_id").(string)
	if from == to {
		return nil
	}
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

func (svc *Service) GetMessages(page, pageSize int, filter string) (messages []*model.Message, count int, err error) {
	message := &model.Message{
		To: svc.ctx.Value("user_id").(string),
	}
	offset := 0
	if page >= 1 {
		offset = (page - 1) * pageSize
	}
	messages, err = message.GetList(svc.db, offset, pageSize, filter)
	if err != nil {
		return
	}
	count, err = message.CountList(svc.db, filter)
	return
}

func (svc *Service) UnreadMessages() (int, error) {
	message := &model.Message{
		To: svc.ctx.Value("user_id").(string),
	}
	return message.CountUnread(svc.db)
}
