package routers

import (
	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestNewRouter(t *testing.T) {
	guard := monkey.Patch(NewRouter, func () *gin.Engine {
		return &gin.Engine{}
	})
	defer guard.Unpatch()
	r := &gin.Engine{}
	require.Equal(t, reflect.TypeOf(NewRouter()).Name(), reflect.TypeOf(r).Name())
}
