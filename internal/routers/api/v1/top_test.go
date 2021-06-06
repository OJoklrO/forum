package v1

import (
	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGetTop(t *testing.T) {
	guard := monkey.Patch(GetTop, func (*gin.Context) {})
	defer guard.Unpatch()
	GetTop(nil)
}
