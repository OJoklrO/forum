package v1

import (
	"github.com/OJoklrO/forum/global"
	"github.com/OJoklrO/forum/internal/service"
	"github.com/OJoklrO/forum/pkg/app"
	"github.com/OJoklrO/forum/pkg/convert"
	"github.com/OJoklrO/forum/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Post struct { }

func NewPost() Post {
	return Post{}
}

// @Summary get post by id
// @Produce json
// @Param id path int true "post id"
// @Success 200 {object} model.Post "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /api/v1/posts/{id} [get]
func (p Post) Get(c *gin.Context) {
	param := service.GetPostRequest{}
	response := app.NewResponse(c)
	//valid, errs := app.BindAndValid(c, &param)
	param.ID = convert.StrTo(c.Param("id")).MustUInt32()
	//if !valid {
	//	global.Logger.Errorf("app.BindAndValid errs: %v", errs)
	//	errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
	//	response.ToErrorResponse(errRsp)
	//	return
	//}
	svc := service.New(c.Request.Context())
	post, err := svc.GetPost(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetPost err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetPostFail)
		return
	}


	response.ToResponse(post)
}

// @Summary get post list
// @Produce json
// @Param page query int false "page number"
// @Param page_size query int false "page size"
// @Param filter query string false "filter"
// @Success 200 {object} model.PostSwagger "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /api/v1/posts [get]
func (p Post) List(c *gin.Context) {
	param := service.GetPostListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BuildAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page: app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	totalRows, err := svc.CountPosts(&service.CountPostsRequest{})
	if err != nil {
		global.Logger.Errorf("svc.CountPosts err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountPostsFail)
		return
	}

	posts, err := svc.GetPostList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetPostList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetPostListFail)
		return
	}
	response.ToResponseList(posts, totalRows)
}

// @Summary create post
// @Produce json
// @Param title body string true "post title"
// @Param desc_img body string true "desc image url"
// @Param createdby body string true "creator"
// @Param content body string true "post content"
// @Success 200 {object} model.Post "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /api/v1/posts [post]
func (p Post) Create(c *gin.Context) {
	param := service.CreatePostRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	calims, err := app.ParseToken(c.Keys["token"].(string))
	if err != nil {
		global.Logger.Errorf("app.ParseToken errs: %v", err)
		response.ToErrorResponse(errcode.NewError(1213123, "developer is sb"))
		return
	}
	param.CreatedBy = calims.Name

	svc := service.New(c.Request.Context())
	pos := svc.CreatePost(&param)
	if pos == nil {
		global.Logger.Errorf("svc.CreatePost err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreatePostFail)
		return
	}

	response.ToResponse(gin.H{
		"post_id": pos.ID,
	})
	return
}

// @Summary delete post need adm cookie
// @Produce json
// @Param id path int true "post id"
// @Success 200 {object} model.Post "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /api/v1/posts/{id} [delete]
func (p Post) Delete(c *gin.Context) {
	param := service.DeletePostRequest{}
	response := app.NewResponse(c)
	param.ID = convert.StrTo(c.Param("id")).MustUInt32()
	//valid, errs := app.BindAndValid(c, &param)
	//if !valid {
	//	global.Logger.Errorf("app.BindAndValid errs: %v", errs)
	//	errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
	//	response.ToErrorResponse(errRsp)
	//	return
	//}

	svc := service.New(c.Request.Context())
	err := svc.DeletePost(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeletePost err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeletePostFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

//func (p Post) Update(c *gin.Context) {
//
//}

func (p Post) CreateTemp(c *gin.Context) {
	param := service.CreatePostRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	//calims, err := app.ParseToken(c.Keys["token"].(string))
	//if err != nil {
	//	global.Logger.Errorf("app.ParseToken errs: %v", err)
	//	response.ToErrorResponse(errcode.NewError(1213123, "developer is sb"))
	//	return
	//}
	//param.CreatedBy = calims.Name

	svc := service.New(c.Request.Context())
	pos := svc.CreatePost(&param)
	if pos == nil {
		global.Logger.Errorf("svc.CreatePost err")
		response.ToErrorResponse(errcode.ErrorCreatePostFail)
		return
	}

	pa := service.CreateCommentRequest{}
	pa.PostID = pos.ID
	pa.Content = param.Content
	pa.CreatedBy = param.CreatedBy

	if err := svc.CreateComment(&pa); err != nil {
		global.Logger.Errorf("svc.CreateComment err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateCommentFail)
		return
	}

	response.ToResponse(gin.H{
		"post_id": pos.ID,
	})
	return
}