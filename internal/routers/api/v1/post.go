package v1

import (
	"forum/internal/model"
	"forum/internal/service"
	"forum/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostHandler struct{}

func NewPost() PostHandler {
	return PostHandler{}
}

type Post struct {
	model.Post
	VoteUp   int `json:"vote_up"`
	VoteDown int `json:"vote_down"`
}

// @Summary Get a post by id
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} Post "Post data"
// @Router /api/v1/posts/{id} [get]
func (p PostHandler) Get(c *gin.Context) {
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

	voteUp, voteDown, err := svc.GetVotes(1, post.ID)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.GetVotes error "+err.Error())
		return
	}

	c.JSON(http.StatusOK, Post{
		VoteUp:   voteUp,
		VoteDown: voteDown,
		Post:     *post,
	})
}

type PostListResponse struct {
	Posts      []Post     `json:"posts"`
	PostBriefs []string   `json:"post_briefs"`
	PostImages [][]string `json:"post_images"`
	TotalPosts int        `json:"total_posts"`
}

// @Summary Get a post list with pagination settings.
// @Produce json
// @Param page query int true "Page number" default(1)
// @Param page_size query int true "Page size" default(20)
// @Param filter query string true "Filter"
// @Param image_mode query bool false "Enable image mode" default(false)
// @Success 200 {object} PostListResponse "success"
// @Router /api/v1/posts [get]
func (p PostHandler) List(c *gin.Context) {
	svc := service.New(c)
	count, err := svc.CountPosts()
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.CountPosts err: "+err.Error())
		return
	}

	page, errPage := strconv.Atoi(c.Query("page"))
	pageSize, errPageSize := strconv.Atoi(c.Query("page_size"))
	pageFilter := c.Query("filter")
	// todo: search filter
	// todo: pin post
	// todo: sort posts
	// todo: posts image mode
	//imageMode := c.Query("image_mode") == "true"
	if errPage != nil || errPageSize != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"page or page_size param error.")
		return
	}

	posts, err := svc.GetPostList(page, pageSize, pageFilter)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.GetPostList err: "+err.Error())
		return
	}

	var respPosts []Post = make([]Post, 0)
	for _, v := range posts {
		voteUp, voteDown, err := svc.GetVotes(1, v.ID)
		if err != nil {
			app.ResponseError(c, http.StatusInternalServerError,
				"svc.GetVotes error "+err.Error())
			return
		}
		respPosts = append(respPosts, Post{
			Post:     *v,
			VoteUp:   voteUp,
			VoteDown: voteDown,
		})
	}

	response := &PostListResponse{Posts: respPosts, TotalPosts: count}
	response.PostBriefs = make([]string, len(respPosts))
	response.PostImages = make([][]string, len(respPosts))
	for i := range respPosts {
		response.PostImages[i], response.PostBriefs[i], err = svc.GetCommentBrief(1, respPosts[i].ID)
		if err != nil {
			app.ResponseError(c, http.StatusInternalServerError,
				"svc.GetCommentBrief: "+err.Error())
		}
	}
	c.JSON(http.StatusOK, response)
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
func (p PostHandler) Create(c *gin.Context) {
	// todo: create a post with emoji, will => have no comment => a post have no comment
	param := service.CreatePostRequest{}

	err := c.ShouldBind(&param)
	if err != nil {
		app.ResponseError(c, http.StatusBadRequest,
			err.Error())
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
func (p PostHandler) Delete(c *gin.Context) {
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
