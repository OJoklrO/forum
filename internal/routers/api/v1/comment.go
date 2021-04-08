package v1

import (
	"github.com/OJoklrO/forum/global"
	"github.com/OJoklrO/forum/internal/service"
	"github.com/OJoklrO/forum/pkg/app"
	"github.com/OJoklrO/forum/pkg/convert"
	"github.com/OJoklrO/forum/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Comment struct {

}

func NewComment() Comment {
	return Comment{}
}

//func (c Comment) Get(ctx *gin.Context) {
//
//}

// @Summary get comment list with post id
// @Produce json
// @Param id path int true "post id"
// @Param page query int false "page number"
// @Param page_size query int false "page size"
// @Success 200 {object} model.CommentSwagger "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /api/v1/comments/{post_id} [get]
func (c Comment) List(ctx *gin.Context) {
	param := service.ListCommentRequest{}
	response := app.NewResponse(ctx)
	param.PostID = convert.StrTo(ctx.Param("post_id")).MustUInt32()

	svc := service.New(ctx.Request.Context())
	pager := app.Pager{
		Page: app.GetPage(ctx),
		PageSize: app.GetPageSize(ctx),
	}
	totalRows, err := svc.CountComments(&service.CountCommentsRequest{PostID: param.PostID})
	if err != nil {
		global.Logger.Errorf("svc.CountComments err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountCommentsFail)
		return
	}

	comments, err := svc.ListComment(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.ListComment err: %v", err)
		response.ToErrorResponse(errcode.ErrorListCommentsFail)
		return
	}
	response.ToResponseList(comments, totalRows)
}

// @Summary create comment
// @Produce json
// @Param content body string true "comment content"
// @Param created_by body int true "creator id"
// @Param post_id body int true "post id"
// @Success 200 {object} model.Comment "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /api/v1/comments/ [post]
func (c Comment) Create(ctx *gin.Context) {
	param := service.CreateCommentRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(ctx.Request.Context())
	err := svc.CreateComment(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateComment err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateCommentFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary delete comment need adm cookie
// @Produce json
// @Param id path int true "comment id"
// @Success 200 {object} model.Comment "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /api/v1/comments/{id} [delete]
func (c Comment) Delete(ctx *gin.Context) {
	param := service.DeleteCommentRequest{}
	param.ID = convert.StrTo(ctx.Param("id")).MustUInt32()
	response := app.NewResponse(ctx)

	svc := service.New(ctx.Request.Context())
	err := svc.DeleteComment(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteComment err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteCommentFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}