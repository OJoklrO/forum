package v1

import (
	"forum/internal/service"
	"forum/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get account information
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} service.AccountInfo "success"
// @Router /api/v1/accounts/{id} [get]
func GetAccountInfo(c *gin.Context) {
	svc := service.New(c)
	account, err := svc.GetUserInfo(c.Params.ByName("id"))
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}

// @Summary Edit account information.
// @Produce json
// @Param body body service.EditAccountInfoRequest true "New account information."
// @Param token header string true "token"
// @Success 200 {object} MessageResponse "success"
// @Router /api/v1/accounts [put]
func EditAccountInfo(c *gin.Context) {
	svc := service.New(c)
	var param service.EditAccountInfoRequest
	err := c.ShouldBind(&param)
	if err != nil {
		app.ResponseError(c, http.StatusBadRequest,
			err.Error())
		return
	}

	err = svc.EditUserInfo(&param)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, MessageResponse{"success."})
}

// @Summary Get account information of myself.
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} service.AccountInfo "success"
// @Router /api/v1/me/account [get]
func GetMyAccountInfo(c *gin.Context) {
	svc := service.New(c)
	id, exists := c.Get("user_id")
	if !exists {
		app.ResponseError(c, http.StatusInternalServerError, "user_id not exists")
		return
	}
	account, err := svc.GetUserInfo(id.(string))
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}

// @Summary Get a comment list by the user id.
// @Produce json
// @Param page query int true "page number" default(1)
// @Param page_size query int true "page size" default(20)
// @Param filter query string false "filter"
// @Param token header string true "token"
// @Success 200 {object} CommentListResponse "success"
// @Router /api/v1/me/comments [get]
func GetMyComments(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	pageSize, errPageSize := strconv.Atoi(c.Query("page_size"))
	filter := c.Query("filter")
	if errPage != nil || errPageSize != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"page or page_size param error.")
		return
	}

	id, exists := c.Get("user_id")
	if !exists {
		app.ResponseError(c, http.StatusInternalServerError, "user_id not exists")
		return
	}

	svc := service.New(c)
	commentCount, err := svc.CountCommentsOfUser(id.(string))
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.CountComments err: "+err.Error())
		return
	}

	comments, err := svc.GetAllCommentsByUser(id.(string), filter, page, pageSize)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.ListComment err: "+err.Error())
		return
	}

	var respComments = make([]Comment, 0)
	for _, v := range comments {
		voteUp, voteDown, voteStatus, err := svc.GetVotes(v.ID, v.PostID)
		if err != nil {
			app.ResponseError(c, http.StatusInternalServerError,
				"svc.GetVotes err: "+err.Error())
			return
		}
		newItem := Comment{
			ID:         v.ID,
			PostID:     v.PostID,
			UserID:     v.UserID,
			Content:    v.Content,
			Time:       app.TimeFormat(v.Time),
			IsEdited:   v.IsEdited,
			VoteUp:     voteUp,
			VoteDown:   voteDown,
			VoteStatus: voteStatus,
		}
		respComments = append(respComments, newItem)
	}
	c.JSON(http.StatusOK, CommentListResponse{
		respComments,
		commentCount,
		0,
	})
}
