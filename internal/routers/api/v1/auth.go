package v1

import (
	"github.com/OJoklrO/forum/global"
	"github.com/OJoklrO/forum/internal/service"
	"github.com/OJoklrO/forum/pkg/app"
	"github.com/OJoklrO/forum/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Token struct {
	Token string
}

// @Summary auth and get token
// @Produce json
// @Param uname query int false "user name"
// @Param upassword query int false "user password"
// @Success 200 {object} Token "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.Uname, param.Upassword)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
		"uname": param.Uname,
	})
}

// @Summary get add user
// @Produce json
// @Param uname query int false "user name"
// @Param upassword query int false "user password"
// @Success 200 {object} model.Auth "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /auth [post]
func CreateAuth(c *gin.Context) {
	param := service.CreateAuthRequest{}
	param.Uname = c.Query("uname")
	param.Upassword = c.Query("upassword")

	response := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	if !svc.AuthExist(&service.AuthExistRequest{Uname: param.Uname}) {
		response.ToErrorResponse(errcode.ErrorAuthExist)
		return
	}

	if err := svc.CreateAuth(&param); err != nil {
		global.Logger.Errorf("svc.CreateAuth err: %v", err)
		response.ToErrorResponse(errcode.ErrorAuthCreateFail)
		return
	}

	response.ToResponse(gin.H{})
}