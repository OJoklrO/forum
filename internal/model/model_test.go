package model

import (
	"bou.ke/monkey"
	"errors"
	"forum/pkg/setting"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDBEngine(t *testing.T) {
	guard := monkey.Patch(NewDBEngine, func (*setting.DatabaseSettings) (*gorm.DB, error) {
		return nil, errors.New("create db error")
	})
	_, err := NewDBEngine(nil)
	require.Error(t, err)
	guard.Unpatch()
	guard = monkey.Patch(NewDBEngine, func (*setting.DatabaseSettings) (*gorm.DB, error) {
		return nil, nil
	})
	_, err = NewDBEngine(nil)
	require.Nil(t, err)
	guard.Unpatch()
}
