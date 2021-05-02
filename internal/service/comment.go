package service

import (
	"forum/internal/model"
	"log"
	"regexp"
	"time"

	"github.com/grokify/html-strip-tags-go"
)

func (svc *Service) CountCommentsOfPost(postId uint32) (int, error) {
	c := model.Comment{
		PostID: postId,
	}
	return c.CountByPostId(svc.db)
}

type ListCommentRequest struct {
	PostID uint32 `form:"post_id" binding:"required,gte=1"`
}

func (svc *Service) CountCommentUsers(param *ListCommentRequest) (int, error) {
	c := model.Comment{
		PostID: param.PostID,
	}
	return c.CountUsers(svc.db)
}

func (svc *Service) ListComment(param *ListCommentRequest, page, pageSize int) ([]*model.Comment, error) {
	c := model.Comment{PostID: param.PostID}
	pageOffset := 0
	if page > 0 {
		pageOffset = (page - 1) * pageSize
	}
	return c.List(svc.db, pageOffset, pageSize)
}

type CreateCommentRequest struct {
	PostID  uint32 `form:"post_id" binding:"gte=0" json:"post_id"`
	Content string `form:"content" binding:"required" json:"content"`
}

func (svc *Service) CreateComment(param *CreateCommentRequest) error {
	comment := model.Comment{
		UserID:  svc.ctx.Value("user_id").(string),
		PostID:  param.PostID,
		Content: param.Content,
		Time:    time.Now().Format("2006-01-02"),
	}
	err := comment.Create(svc.db)
	if err != nil {
		return err
	}

	count, err := svc.CountCommentsOfPost(param.PostID)
	if err != nil {
		return err
	}

	err = svc.UpdateUserLevel()
	if err != nil {
		return err
	}

	post := model.Post{
		ID:           param.PostID,
		LatestReply:  comment.Time,
		ReplyUserID:  svc.ctx.Value("user_id").(string),
		CommentCount: uint32(count),
	}
	err = post.Update(svc.db)
	if err != nil {
		return err
	}

	// find @username
	var matches [][]string
	atRex := regexp.MustCompile(" @(.*?) ")
	matches = atRex.FindAllStringSubmatch(param.Content, -1)
	var names []string
	for _, val := range matches {
		names = append(names, val[1])
		from := val[1]
		to := svc.ctx.Value("user_name").(string)
		err = svc.CreateNotifyMessage(from, to, post.ID, comment.ID)
		if err != nil {
			log.Println("error svc.CreateNotifyMessage: " + err.Error())
		}
	}
	return nil
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
		Time:     time.Now().Format("2006-01-02"),
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

	// find @username
	var matches [][]string
	atRex := regexp.MustCompile(" @(.*?) ")
	matches = atRex.FindAllStringSubmatch(param.Content, -1)
	var names []string
	for _, val := range matches {
		names = append(names, val[1])
		from := val[1]
		to := svc.ctx.Value("user_name").(string)
		err = svc.CreateNotifyMessage(from, to, post.ID, comment.ID)
		if err != nil {
			log.Println("error svc.CreateNotifyMessage: " + err.Error())
		}
	}
	return nil
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

func (svc *Service) Vote(id, postId uint32, support int) error {
	v := &model.Vote{
		UserID:    svc.ctx.Value("user_id").(string),
		PostID:    postId,
		CommentID: id,
		Vote:      support,
	}

	return v.SetOrCreate(svc.db)
}

func (svc *Service) GetVotes(id, postId uint32) (int, int, error) {
	v := &model.Vote{
		PostID:    postId,
		CommentID: id,
	}
	return v.CommentVoteCount(svc.db)
}

func getBrief(content string) (imageURLs []string, result string) {
	var matches [][]string
	imageRex := regexp.MustCompile("<img.*?src=\"(.*?)\"(.*?)alt=\"(.*?)\"> ")
	matches = imageRex.FindAllStringSubmatch(content, -1)
	cleanedContent := imageRex.ReplaceAllString(content, "")

	for _, val := range matches {
		imageURLs = append(imageURLs, val[1])
	}

	result = strip.StripTags(cleanedContent)
	return
}

func (svc *Service) GetCommentBrief(id, postId uint32) (imageURLs []string, briefContent string, err error) {
	comment, err := svc.GetComment(id, postId)
	if err != nil {
		return
	}
	imageURLs, briefContent = getBrief(comment.Content)
	return
}
