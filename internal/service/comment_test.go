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

func TestService_CountAllCommentsOfUser(t *testing.T) {
	c := &model.Comment{}
	svc := New(nil)
	// step1: count by user error
	guard := monkey.PatchInstanceMethod(reflect.TypeOf(c), "CountByUser",
		func (*model.Comment, *gorm.DB) (int, error) {
			return 0, errors.New("count by user error")
		})
	_, err := svc.CountAllCommentsOfUser("any")
	require.Error(t, err)
	guard.Unpatch()
	// step2: count by user success
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(c), "CountByUser",
		func (*model.Comment, *gorm.DB) (int, error) {
			return 0, nil
		})
	count, err := svc.CountAllCommentsOfUser("any")
	require.Nil(t, err)
	require.Equal(t, count, 0)
	guard.Unpatch()
}

func TestService_CountCommentsOfPost(t *testing.T) {
	c := &model.Comment{}
	svc := New(nil)
	// step1: count by postId error
	guard := monkey.PatchInstanceMethod(reflect.TypeOf(c), "CountByPostId",
		func (*model.Comment, *gorm.DB) (int, error) {
			return 0, errors.New("count by postId error")
		})
	_, err := svc.CountCommentsOfPost(0)
	require.Error(t, err)
	guard.Unpatch()
	// step2: count by postId success
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(c), "CountByPostId",
		func (*model.Comment, *gorm.DB) (int, error) {
			return 0, nil
		})
	count, err := svc.CountCommentsOfPost(0)
	require.Nil(t, err)
	require.Equal(t, count, 0)
	guard.Unpatch()
}

func TestService_CountCommentUsers(t *testing.T) {
	c := &model.Comment{}
	svc := New(nil)
	// step1: count by postId error
	guard := monkey.PatchInstanceMethod(reflect.TypeOf(c), "CountUsers",
		func (*model.Comment, *gorm.DB) (int, error) {
			return 0, errors.New("count by postId error")
		})
	_, err := svc.CountCommentUsers(0)
	require.Error(t, err)
	guard.Unpatch()
	// step2: count by postId success
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(c), "CountUsers",
		func (*model.Comment, *gorm.DB) (int, error) {
			return 0, nil
		})
	count, err := svc.CountCommentUsers(0)
	require.Nil(t, err)
	require.Equal(t, count, 0)
	guard.Unpatch()
}

func TestService_ListComment(t *testing.T) {
	c := &model.Comment{}
	svc := New(nil)
	// step1: model.Comment.List error
	guard := monkey.PatchInstanceMethod(reflect.TypeOf(c), "List",
		func (*model.Comment, *gorm.DB, int, int, string) ([]*model.Comment, error) {
			return nil, errors.New("count by postId error")
		})
	_, err := svc.ListComment(0, 0, 0, "any")
	require.Error(t, err)
	guard.Unpatch()
	// step2: model.Comment.List success
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(c), "List",
		func (*model.Comment, *gorm.DB, int, int, string) ([]*model.Comment, error) {
			return nil, nil
		})
	_, err = svc.ListComment(0, 0, 0, "any")
	require.Nil(t, err)
	guard.Unpatch()
}

func TestService_GetComment(t *testing.T) {
	c := &model.Comment{}
	svc := New(nil)
	guard := monkey.PatchInstanceMethod(reflect.TypeOf(*c), "Get",
		func (model.Comment, *gorm.DB) (*model.Comment, error) {
			return nil, errors.New("model.Comment.Get error")
		})
	defer guard.Unpatch()
	_, err := svc.GetComment(0, 0)
	require.Error(t, err)
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(*c), "Get",
		func (model.Comment, *gorm.DB) (*model.Comment, error) {
			return nil, nil
		})
	defer guard.Unpatch()
	_, err = svc.GetComment(0, 0)
	require.Nil(t, err)
}
