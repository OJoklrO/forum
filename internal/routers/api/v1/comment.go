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

type Comment struct{}

func NewComment() Comment {
	return Comment{}
}

type CommentListResponse struct {
	Comments   []*model.Comment `json:"comments"`
	TotalPages int              `json:"total_pages"`
}

// @Summary Get a comment list by the post id.
// @Produce json
// @Param post_id path int true "post id"
// @Param page query int true "page number" default(1)
// @Param page_size query int true "page size" default(20)
// @Success 200 {object} CommentListResponse "success"
// @Router /api/v1/comments/{post_id} [get]
func (comment *Comment) List(c *gin.Context) {
	param := service.ListCommentRequest{}
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		app.ResponseError(c, http.StatusBadRequest, "param error.")
		return
	}
	param.PostID = uint32(postID)

	svc := service.New(c)
	totalPages, err := svc.CountComments(&service.ListCommentRequest{PostID: param.PostID})
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.CountComments err: "+err.Error())
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
	c.JSON(http.StatusOK, CommentListResponse{
		comments,
		totalPages,
	})
}

// @Summary Comment a post.
// @Produce json
// @Param body body service.CreateCommentRequest true "body"
// @Param token header string true "jwt token"
// @Success 200 {object} MessageResponse "success"
// @Router /api/v1/comments [post]
func (comment *Comment) Create(c *gin.Context) {
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
func (comment *Comment) Delete(c *gin.Context) {
	param := service.LocateCommentRequest{}
	id, errId := strconv.Atoi(c.Param("id"))
	postId, errPost := strconv.Atoi(c.Param("post_id"))
	if errId != nil || errPost != nil {
		app.ResponseError(c, http.StatusBadRequest, "param error.")
		return
	}
	param.ID = uint32(id)
	param.PostID = uint32(postId)

	svc := service.New(c)
	err := svc.DeleteComment(&param)
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
func (comment *Comment) Edit(c *gin.Context) {
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

// @Summary Delete a comment.
// @Produce json
// @Param post_id path int true "post id"
// @Param id path int true "comment id"
// @Success 200 {object} model.Comment "success"
// @Router /api/v1/comments/{post_id}/{id} [get]
func (comment *Comment) Get(c *gin.Context) {
	param := service.LocateCommentRequest{}
	id, errId := strconv.Atoi(c.Param("id"))
	postId, errPost := strconv.Atoi(c.Param("post_id"))
	if errId != nil || errPost != nil {
		app.ResponseError(c, http.StatusBadRequest, "param error.")
		return
	}
	param.ID = uint32(id)
	param.PostID = uint32(postId)
	svc := service.New(c)
	targetComment, err := svc.GetComment(&param)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.GetComment err: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, targetComment)
}
