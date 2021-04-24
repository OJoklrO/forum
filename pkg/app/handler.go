package app

import (
	"github.com/OJoklrO/forum/global"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ResponseError(c *gin.Context, code int, msg string) {
	global.Logger.Error(msg)
	c.JSON(code, gin.H{
		"msg": msg,
	})
	return
}

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// BindBodyWithValidation binds values from POST request body.
func BindBodyWithValidation(c *gin.Context, v interface{}) ValidErrors {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		vErrs, _ := err.(validator.ValidationErrors)
		for key, value := range vErrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return errs
	}
	return nil
}
