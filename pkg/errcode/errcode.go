package errcode

import (
	"fmt"
	"net/http"
)

// todo: remove err code
var (
	ErrorAuthExist      = NewError(30000001, "user has exist")
	ErrorAuthCreateFail = NewError(30000002, "create user error")
)

var (
	Success                   = NewError(0, "success")
	ServerError               = NewError(10000000, "server error")
	InvalidParams             = NewError(10000001, "invalid params")
	NotFound                  = NewError(10000002, "not found")
	UnauthorizedAuthNotExist  = NewError(10000003, "auth not exist")
	UnauthorizedTokenError    = NewError(10000004, "token error")
	UnauthorizedTokenTimeout  = NewError(10000005, "token time out")
	UnauthorizedTokenGenerate = NewError(10000006, "token generate error")
	TooManyRequests           = NewError(10000007, "too many request")
)

var (
	ErrorCountPostsFail  = NewError(20010001, "count posts error")
	ErrorGetPostFail     = NewError(20010002, "get post error")
	ErrorGetPostListFail = NewError(20010003, "get post list error")
	ErrorCreatePostFail  = NewError(20010004, "create post error")
	ErrorDeletePostFail  = NewError(20010005, "delete post error")

	ErrorListCommentsFail  = NewError(20020001, "list comments error")
	ErrorCountCommentsFail = NewError(20020002, "count comments error")
	ErrorCreateCommentFail = NewError(20020003, "create comment error")
	ErrorDeleteCommentFail = NewError(20020004, "delete comment error")
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("error code %d is already exist, please change it", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, info: %v", e.code, e.msg)

}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.code {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
