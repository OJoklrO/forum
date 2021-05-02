package v1

import (
	"forum/internal/service"
	"forum/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
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

// todo: get user's comments and posts
