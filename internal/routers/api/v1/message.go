package v1

import (
	"forum/internal/model"
	"forum/internal/service"
	"forum/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MessageListResponse struct {
	Messages []*model.Message `json:"messages"`
	Count    int              `json:"count"`
}

// @Summary Get messages(reply notifications).
// @Produce json
// @Param token header string true "jwt token"
// @Param filter query string false "Filter"
// @Param page query int true "Page number" default(1)
// @Param page_size query int true "Page size" default(20)
// @Success 200 {object} MessageListResponse "messages"
// @Router /api/v1/messages [get]
func GetMessageList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	pageSize, errPageSize := strconv.Atoi(c.Query("page_size"))
	pageFilter := c.Query("filter")
	if errPage != nil || errPageSize != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"page or page_size param error.")
		return
	}
	svc := service.New(c)
	results, count, err := svc.GetMessages(page, pageSize, pageFilter)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.GetUnreadMessages: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, MessageListResponse{Messages: results, Count: count})
}

type UnreadMessageCountResponse struct {
	Count int `json:"count"`
}

// @Summary Get the number of unread messages(reply notifications).
// @Produce json
// @Param token header string true "jwt token"
// @Success 200 {object} UnreadMessageCountResponse "unread message count"
// @Router /api/v1/messages/unread [get]
func GetUnreadMessageCount(c *gin.Context) {
	svc := service.New(c)
	result, err := svc.UnreadMessages()
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.GetUnreadMessages: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, UnreadMessageCountResponse{result})
}

// @Summary Make an unread message read.
// @Produce json
// @Param token header string true "jwt token"
// @Param post_id path int true "post id"
// @Param comment_id path int true "comment id"
// @Success 200 {object} MessageResponse "success"
// @Router /api/v1/messages/{post_id}/{comment_id} [get]
func ReadMessage(c *gin.Context) {
	postId, errPostParam := strconv.Atoi(c.Param("post_id"))
	commentId, errCommentParam := strconv.Atoi(c.Param("comment_id"))
	if errPostParam != nil || errCommentParam != nil {
		app.ResponseError(c, http.StatusBadRequest, "param error")
	}

	svc := service.New(c)
	err := svc.ReadMessage(c.Value("user_id").(string),
		uint32(postId), uint32(commentId))
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.GetUnreadMessages: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, MessageResponse{"success."})
}
