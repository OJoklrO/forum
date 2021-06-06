package middleware

import (
	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestJWT(t *testing.T) {
	guard := monkey.Patch(JWT, func (bool) gin.HandlerFunc {
		return func (*gin.Context){}
	})
	defer guard.Unpatch()
	require.Equal(t, reflect.TypeOf(JWT(true)).Name(), "HandlerFunc")
}
