package service

import (
	"forum/internal/model"
	"forum/pkg/app"
	"log"
	"regexp"
	"time"
)

func (svc *Service) CountAllCommentsOfUser(userId string) (int, error) {
	c := model.Comment{
		UserID: userId,
	}
	return c.CountByUser(svc.db)
}

func (svc *Service) CountCommentsOfPost(postId uint32) (int, error) {
	c := model.Comment{
		PostID: postId,
	}
	return c.CountByPostId(svc.db)
}

func (svc *Service) CountCommentUsers(postId uint32) (int, error) {
	c := model.Comment{
		PostID: postId,
	}
	return c.CountUsers(svc.db)
}

func (svc *Service) ListComment(postId uint32, page, pageSize int, filter string) ([]*model.Comment, error) {
	c := model.Comment{PostID: postId}
	pageOffset := 0
	if page > 0 {
		pageOffset = (page - 1) * pageSize
	}
	return c.List(svc.db, pageOffset, pageSize, filter)
}

type CreateCommentRequest struct {
	PostID  uint32 `form:"post_id" binding:"gte=0" json:"post_id"`
	Content string `form:"content" binding:"required" json:"content"`
}

func (svc *Service) CreateComment(param *CreateCommentRequest) (*model.Comment, error) {
	comment := model.Comment{
		UserID:  svc.ctx.Value("user_id").(string),
		PostID:  param.PostID,
		Content: param.Content,
		Time:    time.Now().Unix(),
	}
	err := comment.Create(svc.db)
	if err != nil {
		return nil, err
	}

	count, err := svc.CountCommentsOfPost(param.PostID)
	if err != nil {
		return nil, err
	}

	err = svc.UpdateUserLevel()
	if err != nil {
		return nil, err
	}

	post := model.Post{
		ID:           param.PostID,
		LatestReply:  comment.Time,
		ReplyUserID:  svc.ctx.Value("user_id").(string),
		CommentCount: uint32(count),
	}
	err = post.Update(svc.db)
	if err != nil {
		return nil, err
	}

	names := getNamesFromContent(param.Content)
	for _, name := range names {
		to := name
		err = svc.CreateNotifyMessage(to, post.ID, comment.ID)
		if err != nil {
			log.Println("error svc.CreateNotifyMessage: " + err.Error())
		}
	}
	return &comment, nil
}

type EditCommentRequest struct {
	PostID  uint32 `form:"post_id" binding:"gte=0" json:"post_id"`
	ID      uint32 `form:"id" binding:"gte=0" json:"id"`
	Content string `form:"content" binding:"required" json:"content"`
}

func (svc *Service) EditComment(param *EditCommentRequest) error {
	comment := model.Comment{
		PostID:   param.PostID,
		ID:       param.ID,
		Content:  param.Content,
		Time:     time.Now().Unix(),
		IsEdited: true,
	}
	err := comment.Update(svc.db)
	if err != nil {
		return err
	}
	post := model.Post{
		ID:          param.PostID,
		LatestReply: comment.Time,
		ReplyUserID: svc.ctx.Value("user_id").(string),
	}
	err = post.Update(svc.db)
	if err != nil {
		return err
	}

	names := getNamesFromContent(param.Content)
	for _, name := range names {
		to := name
		err = svc.CreateNotifyMessage(to, post.ID, comment.ID)
		if err != nil {
			log.Println("error svc.CreateNotifyMessage: " + err.Error())
		}
	}
	return nil
}

func getNamesFromContent(content string) []string {
	if len(content) == 0 {
		return nil
	}

	_, content = app.CleanHTMLTags(content)
	content += " " // add this for @userID in the last

	var names []string

	// find first @userid at the beginning
	if content[0] == '@' {
		index := 0
		for ; index < len(content); index++ {
			if content[index] == ' ' {
				names = append(names, content[1:index])
				content = content[index+1:]
				break
			}
		}
		// if this can not found,
		// there is nothing after @username
		if len(names) == 0 {
			names = append(names, content[1:])
			return names
		}
	}

	var matches [][]string
	atRex := regexp.MustCompile(" @(.*?) ")
	matches = atRex.FindAllStringSubmatch(content, -1)

	for _, val := range matches {
		names = append(names, val[1])
	}
	return names
}

func (svc *Service) DeleteComment(id, postId uint32) error {
	comment, err := svc.GetComment(id, postId)
	if err != nil {
		return err
	}
	err = svc.CheckCommonPermission(comment.UserID)
	if err != nil {
		return err
	}
	return comment.Delete(svc.db)
}

func (svc *Service) GetComment(id, postId uint32) (*model.Comment, error) {
	c := model.Comment{
		ID:     id,
		PostID: postId,
	}
	return c.Get(svc.db)
}

func (svc *Service) GetAllCommentsByUser(userId, filter string, page, pageSize int) ([]*model.Comment, error) {
	c := model.Comment{UserID: userId}
	pageOffset := 0
	if page > 0 {
		pageOffset = (page - 1) * pageSize
	}
	return c.List(svc.db, pageOffset, pageSize, filter)
}

func (svc *Service) Vote(id, postId uint32, support int) error {
	v := &model.Vote{
		UserID:    svc.ctx.Value("user_id").(string),
		PostID:    postId,
		CommentID: id,
		Vote:      support,
	}

	return v.SetOrCreate(svc.db)
}

func (svc *Service) GetVotes(id, postId uint32) (up int, down int, status int, err error) {
	v := &model.Vote{
		PostID:    postId,
		CommentID: id,
	}
	up, down, err = v.CommentVoteCount(svc.db)
	if err != nil {
		return
	}

	userId, exist := svc.ctx.Get("user_id")
	if exist {
		v.UserID = userId.(string)
		errVoteStatus := v.Get(svc.db)
		if errVoteStatus != nil {
			return // status is still 0
		}
		status = v.Vote
	}
	return
}

func (svc *Service) GetCommentBrief(id, postId uint32) (imageURLs []string, briefContent string, err error) {
	comment, err := svc.GetComment(id, postId)
	if err != nil {
		return
	}
	imageURLs, briefContent = app.CleanHTMLTags(comment.Content)
	return
}
