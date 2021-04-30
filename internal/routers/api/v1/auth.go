package v1

import (
	"forum/global"
	"forum/internal/service"
	"forum/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type LoginResponse struct {
	Token string `json:"token"`
}

// @Summary Log in.
// @Produce json
// @Param body body service.LoginRequest true "body"
// @Success 200 {object} LoginResponse "success"
// @Router /api/v1/accounts/login [post]
func Login(c *gin.Context) {
	param := service.LoginRequest{}
	errs := app.BindBodyWithValidation(c, &param)
	if errs != nil {
		app.ResponseError(c, http.StatusBadRequest,
			"app.BindBodyWithValidation errs: "+strings.Join(errs.Errors(), ","))
		return
	}

	svc := service.New(c)
	token, err := svc.LoginAccount(&param)
	if err != nil {
		app.ResponseError(c, http.StatusUnauthorized,
			"svc.LoginAccount err: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, LoginResponse{token})
}

// @Summary Register. The invite code is "xduxdu".
// @Produce json
// @Param body body service.RegisterRequest true "body"
// @Success 200 {object} LoginResponse "success"
// @Router /api/v1/accounts/register [post]
func Register(c *gin.Context) {
	param := service.RegisterRequest{}
	errors := app.BindBodyWithValidation(c, &param)
	if errors != nil {
		app.ResponseError(c, http.StatusBadRequest,
			"Param errors: "+strings.Join(errors.Errors(), ", "))
		return
	}

	// for now, there is only one invite code
	if param.InviteCode != global.AppSetting.InviteCode {
		app.ResponseError(c, http.StatusForbidden,
			"The invite code is not valid")
		return
	}

	svc := service.New(c)
	token, err := svc.RegisterAccount(&param)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.Register: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, LoginResponse{token})
}

// @Summary Delete an account.
// @Produce json
// @Param id path string true "id"
// @Param token header string true "jwt token"
// @Success 200 {object} MessageResponse "success"
// @Router /api/v1/accounts/{id} [delete]
func DeleteAccount(c *gin.Context) {
	param := service.DeleteAccountRequest{}
	param.ID = c.Param("id")
	svc := service.New(c)
	err := svc.DeleteAccount(&param)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.DeleteAccount error:"+err.Error())
		return
	}
	c.JSON(http.StatusOK, MessageResponse{"success."})
}
