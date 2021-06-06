package service

import (
	"bou.ke/monkey"
	"errors"
	"forum/internal/model"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestService_GetCheckInRecords(t *testing.T) {
	check := &model.Checkin{}
	svc := New(nil)
	// step1: check.GetThisMonth error
	guard := monkey.PatchInstanceMethod(reflect.TypeOf(check), "GetThisMonth",
		func(*model.Checkin, *gorm.DB) ([]model.Checkin, error){
		return nil, errors.New("GetThisMonth error")
	})
	_, err := svc.GetCheckInRecords("any")
	require.Error(t, err)
	guard.Unpatch()

	// step2: get records succeed
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(check), "GetThisMonth",
		func(*model.Checkin, *gorm.DB) ([]model.Checkin, error){
			return nil, nil
		})
	_, err = svc.GetCheckInRecords("any")
	require.Nil(t, err)
	guard.Unpatch()
}
