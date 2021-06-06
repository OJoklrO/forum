package logger

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/require"
	"io"
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	guard := monkey.Patch(NewLogger, func (io.Writer, string, int) *Logger {
		return &Logger{}
	})
	l := &Logger{}
	defer guard.Unpatch()
	require.Equal(t, reflect.TypeOf(NewLogger(nil, "any", 0)).Name(), reflect.TypeOf(l).Name())
}
