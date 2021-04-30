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
	// todo: reply comments counter in resoponse
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"ID param err: "+err.Error())
		return
	}

	param := service.GetPostRequest{}
	param.ID = uint32(id)
	svc := service.New(c)
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
	TotalPosts int           `json:"total_posts"`
}

// @Summary Get a post list with pagination settings.
// @Produce json
// @Param page query int true "Page number" default(1)
// @Param page_size query int true "Page size" default(20)
// @Success 200 {object} PostListResponse "success"
// @Router /api/v1/posts [get]
func (p Post) List(c *gin.Context) {
	// todo: date(to comments)
	svc := service.New(c)
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

// @Summary Create a post.
// @Produce json
// @Param body body service.CreatePostRequest true "body"
// @Param token header string true "jwt token"
// @Success 200 {object} PostCreateResponse "success"
// @Router /api/v1/posts [post]
func (p Post) Create(c *gin.Context) {
	param := service.CreatePostRequest{}
	errs := app.BindBodyWithValidation(c, &param)
	if errs != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"app.BindBodyWithValidation errs: "+strings.Join(errs.Errors(), ","))
		return
	}

	svc := service.New(c)
	post, err := svc.CreatePost(&param)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.CreatePost error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, PostCreateResponse{
		post.ID,
	})
}

// @Summary Delete a post.
// @Produce json
// @Param id path int true "post id"
// @Param token header string true "jwt token"
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

	svc := service.New(c)
	err = svc.DeletePost(&param)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.DeletePost error.")
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

type PostImageListResponse struct {
	Posts      []*model.Post `json:"posts"`
	Images     []string      `json:"images"`
	TotalPages int           `json:"total_pages"`
}

// @Summary Get a post image list with pagination settings.
// @Produce json
// @Param page query int true "Page number" default(1)
// @Param page_size query int true "Page size" default(20)
// @Success 200 {object} PostImageListResponse "success"
// @Router /api/v1/posts_images [get]
func (p Post) Images(c *gin.Context) {
}
