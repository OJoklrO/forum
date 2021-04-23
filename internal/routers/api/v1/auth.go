package v1

import (
	"github.com/OJoklrO/forum/internal/service"
	"github.com/OJoklrO/forum/pkg/app"
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
// @Router /account/login [post]
func Login(c *gin.Context) {
	param := service.LoginRequest{}
	errs := app.BindBodyWithValidation(c, &param)
	if errs != nil {
		app.ResponseError(c, http.StatusBadRequest,
			"app.BindBodyWithValidation errs: "+strings.Join(errs.Errors(), ","))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.LoginAccount(&param)
	if err != nil {
		app.ResponseError(c, http.StatusUnauthorized,
			"svc.LoginAccount err: "+err.Error())
		return
	}

	token, err := app.GenerateJWTToken(param.ID, param.Password)
	if err != nil {
		app.ResponseError(c, http.StatusUnauthorized,
			"app.GenerateJWTToken err: %v"+err.Error())
		return
	}

	c.JSON(http.StatusOK, LoginResponse{token})
}

// @Summary Register
// @Produce json
// @Param body body service.RegisterRequest true "body"
// @Success 200 {object} MessageResponse "success"
// @Router /account/register [post]
func Register(c *gin.Context) {
	param := service.RegisterRequest{}
	errors := app.BindBodyWithValidation(c, &param)
	if errors != nil {
		app.ResponseError(c, http.StatusBadRequest,
			"Param errors: "+strings.Join(errors.Errors(), ", "))
	}

	svc := service.New(c.Request.Context())
	if err := svc.RegisterAccount(&param); err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.Register: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, MessageResponse{"success."})
}

// @Summary Delete account
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} MessageResponse "success"
// @Router /account/delete/{id} [delete]
func DeleteAccount(c *gin.Context) {
	param := service.DeleteAccountRequest{}
	param.ID = c.Param("id")
	svc := service.New(c.Request.Context())
	err := svc.DeleteAccount(&param)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.DeleteAccount error:"+err.Error())
		return
	}
	c.JSON(http.StatusOK, MessageResponse{"success."})
}
