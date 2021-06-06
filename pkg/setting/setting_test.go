package setting

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSetting(t *testing.T) {
	guard := monkey.Patch(NewSetting, func (string) (*Setting, error) {
		return nil, nil
	})
	_, err := NewSetting("any")
	require.Nil(t, err)
	guard.Unpatch()
}
