package v1

import (
	"forum/internal/model"
	"forum/internal/service"
	"forum/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type CommentHandler struct{}

func NewComment() CommentHandler {
	return CommentHandler{}
}

type Comment struct {
	model.Comment
	Vote int `json:"vote"`
}

type CommentListResponse struct {
	Comments     []Comment `json:"comments"`
	CommentCount int       `json:"comment_count"`
	UserCount    int       `json:"user_count"`
}

// @Summary Get a comment list by the post id.
// @Produce json
// @Param post_id path int true "post id"
// @Param page query int true "page number" default(1)
// @Param page_size query int true "page size" default(20)
// @Success 200 {object} CommentListResponse "success"
// @Router /api/v1/comments/{post_id} [get]
func (comment *CommentHandler) List(c *gin.Context) {
	// todo: return like status depending on jwt (user)
	param := service.ListCommentRequest{}
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		app.ResponseError(c, http.StatusBadRequest, "param error.")
		return
	}
	param.PostID = uint32(postID)

	svc := service.New(c)
	commentCount, err := svc.CountCommentsOfPost(param.PostID)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.CountComments err: "+err.Error())
		return
	}
	userCount, err := svc.CountCommentUsers(&service.ListCommentRequest{PostID: param.PostID})
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.CountCommentUsers err: "+err.Error())
		return
	}

	page, errPage := strconv.Atoi(c.Query("page"))
	pageSize, errPageSize := strconv.Atoi(c.Query("page_size"))
	if errPage != nil || errPageSize != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"page or page_size param error.")
		return
	}

	comments, err := svc.ListComment(&param, page, pageSize)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.ListComment err: "+err.Error())
		return
	}

	var respComments []Comment
	for _, v := range comments {
		vote, err := svc.GetVotes(v.ID, v.PostID)
		if err != nil {
			app.ResponseError(c, http.StatusInternalServerError,
				"svc.GetVotes err: "+err.Error())
			return
		}
		newItem := Comment{
			Vote:    vote,
			Comment: *v,
		}
		respComments = append(respComments, newItem)
	}
	c.JSON(http.StatusOK, CommentListResponse{
		respComments,
		commentCount,
		userCount,
	})
}

// @Summary Comment a post.
// @Produce json
// @Param body body service.CreateCommentRequest true "body"
// @Param token header string true "jwt token"
// @Success 200 {object} MessageResponse "success"
// @Router /api/v1/comments [post]
func (comment *CommentHandler) Create(c *gin.Context) {
	param := service.CreateCommentRequest{}
	errs := app.BindBodyWithValidation(c, &param)
	if errs != nil {
		app.ResponseError(c, http.StatusBadRequest,
			"BindBodyWithValidation errs: "+strings.Join(errs.Errors(), ", "))
		return
	}

	svc := service.New(c)
	err := svc.CreateComment(&param)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.CreateComment err: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, MessageResponse{"success."})
}

// @Summary Delete a comment.
// @Produce json
// @Param post_id path int true "post id"
// @Param id path int true "comment id"
// @Param token header string true "jwt token"
// @Success 200 {object} MessageResponse "success"
// @Router /api/v1/comments/{post_id}/{id} [delete]
func (comment *CommentHandler) Delete(c *gin.Context) {
	id, errId := strconv.Atoi(c.Param("id"))
	postId, errPost := strconv.Atoi(c.Param("post_id"))
	if errId != nil || errPost != nil {
		app.ResponseError(c, http.StatusBadRequest, "param error.")
		return
	}
	svc := service.New(c)
	err := svc.DeleteComment(uint32(id), uint32(postId))
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.DeleteComment err: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, MessageResponse{"success."})
}

// @Summary Edit a post.
// @Produce json
// @Param body body service.EditCommentRequest true "body"
// @Param token header string true "jwt token"
// @Success 200 {object} MessageResponse "success"
// @Router /api/v1/comments [put]
func (comment *CommentHandler) Edit(c *gin.Context) {
	param := service.EditCommentRequest{}
	errs := app.BindBodyWithValidation(c, &param)
	if errs != nil {
		app.ResponseError(c, http.StatusBadRequest,
			"BindBodyWithValidation errs: "+strings.Join(errs.Errors(), ", "))
		return
	}

	svc := service.New(c)
	err := svc.EditComment(&param)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.EditComment err: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, MessageResponse{"success."})
}

// @Summary Get a single comment.
// @Produce json
// @Param post_id path int true "post id"
// @Param id path int true "comment id"
// @Success 200 {object} model.Comment "success"
// @Router /api/v1/comments/{post_id}/{id} [get]
func (comment *CommentHandler) Get(c *gin.Context) {
	id, errId := strconv.Atoi(c.Param("id"))
	postId, errPost := strconv.Atoi(c.Param("post_id"))
	if errId != nil || errPost != nil {
		app.ResponseError(c, http.StatusBadRequest, "param error.")
		return
	}
	svc := service.New(c)
	targetComment, err := svc.GetComment(uint32(id), uint32(postId))
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.GetComment err: "+err.Error())
		return
	}

	vote, err := svc.GetVotes(uint32(id), uint32(postId))
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.GetVotes err: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, Comment{
		Comment: *targetComment,
		Vote:    vote,
	})
}

// @Summary Vote on a comment.
// @Produce json
// @Param post_id path int true "post id"
// @Param id path int true "comment id"
// @Param support path int true "-1 or 1"
// @Param token header string true "jwt token"
// @Success 200 {object} MessageResponse "success"
// @Router /api/v1/comments/{post_id}/{id}/vote/{support} [get]
func (comment *CommentHandler) Vote(c *gin.Context) {
	id, errId := strconv.Atoi(c.Param("id"))
	postId, errPost := strconv.Atoi(c.Param("post_id"))
	supportValue, errVote := strconv.Atoi(c.Param("support"))
	if errId != nil || errPost != nil || errVote != nil {
		app.ResponseError(c, http.StatusBadRequest, "param error.")
		return
	}

	if !(supportValue == -1 || supportValue == 1) {
		app.ResponseError(c, http.StatusBadRequest, "vote 'support' param should be -1 or 1.")
		return
	}

	svc := service.New(c)
	err := svc.Vote(uint32(id), uint32(postId), supportValue)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.Vote err: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, MessageResponse{"success."})
}
