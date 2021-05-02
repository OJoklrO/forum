package app

import (
	"forum/global"
	"github.com/gin-gonic/gin"
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
	// todo: validator, middleware of translation
	_ = c.ShouldBind(v)
	return nil
	//var errs ValidErrors
	//err := c.ShouldBind(v)
	//if err != nil {
	//	v := c.Value("trans")
	//	trans, _ := v.(ut.Translator)
	//	vErrs, _ := err.(validator.ValidationErrors)
	//	for key, value := range vErrs.Translate(trans) {
	//		errs = append(errs, &ValidError{
	//			Key:     key,
	//			Message: value,
	//		})
	//	}
	//	return errs
	//}
	//return nil
}
