package v1

import (
	"github.com/OJoklrO/forum/internal/model"
	"github.com/OJoklrO/forum/internal/service"
	"github.com/OJoklrO/forum/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Post struct{}

func NewPost() Post {
	return Post{}
}

// @Summary Get a post by id
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} model.Post "Post data"
// @Router /api/v1/posts/{id} [get]
func (p Post) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"ID param err: "+err.Error())
		return
	}

	param := service.GetPostRequest{}
	param.ID = uint32(id)
	svc := service.New(c.Request.Context())
	post, err := svc.GetPost(&param)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.Get error "+err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}

type PostListResponse struct {
	Posts      []*model.Post `json:"posts"`
	TotalPages int           `json:"total_pages"`
}

// @Summary Get a post list with pagination settings.
// @Produce json
// @Param page query int true "Page number"
// @Param page_size query int true "Page size"
// @Success 200 {object} PostListResponse "success"
// @Router /api/v1/posts [get]
func (p Post) List(c *gin.Context) {
	svc := service.New(c.Request.Context())
	count, err := svc.CountPosts()
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.CountPosts err: "+err.Error())
		return
	}

	page, errPage := strconv.Atoi(c.Query("page"))
	pageSize, errPageSize := strconv.Atoi(c.Query("page_size"))
	if errPage != nil || errPageSize != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"page or page_size param error.")
		return
	}

	posts, err := svc.GetPostList(page, pageSize)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.GetPostList err: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, PostListResponse{
		posts,
		count,
	})
}

type PostCreateResponse struct {
	PostID uint32 `json:"post_id"`
}

// @Summary create post
// @Produce json
// @Param body body service.CreatePostRequest true "body"
// @Success 200 {object} PostCreateResponse "success"
// @Router /api/v1/posts [post]
func (p Post) Create(c *gin.Context) {
	param := service.CreatePostRequest{}
	// todo: validator (no value check; no require check for json)
	errs := app.BindBodyWithValidation(c, &param)
	if errs != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"app.BindBodyWithValidation errs: "+strings.Join(errs.Errors(), ","))
		return
	}

	// todo: jwt
	//calims, err := app.ParseToken(c.Keys["token"].(string))
	//if err != nil {
	//	global.Logger.Errorf("app.ParseToken errs: %v", err)
	//	response.ToErrorResponse(errcode.NewError(1213123, "developer is sb"))
	//	return
	//}
	//param.UserID = convert.StrTo(calims.ID).MustUInt32()

	// todo: apply content to comments(create post)

	svc := service.New(c.Request.Context())
	post := svc.CreatePost(&param)
	if post == nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.CreatePost err")
		return
	}

	c.JSON(http.StatusOK, PostCreateResponse{
		post.ID,
	})
}

// @Summary Delete a post.
// @Produce json
// @Param id path int true "post id"
// @Success 200 {object} model.Post "success"
// @Router /api/v1/posts/{id} [delete]
func (p Post) Delete(c *gin.Context) {
	param := service.DeletePostRequest{}
	id, err := strconv.Atoi(c.Param("id"))
	param.ID = uint32(id)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"id param error.")
		return
	}

	svc := service.New(c.Request.Context())
	err = svc.DeletePost(&param)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.DeletePost error.")
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
