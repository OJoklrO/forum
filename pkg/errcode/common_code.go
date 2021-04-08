package errcode

var (
	Success                 = NewError(0, "success")
	ServerError             = NewError(10000000, "server error")
	InvalidParams           = NewError(10000001, "invalid params")
	NotFound                = NewError(10000002, "not found")
	UnauthorizedAuthNotExist= NewError(10000003, "auth not exist")
	UnauthorizedTokenError  = NewError(10000004, "token error")
	UnauthorizedTokenTimeout= NewError(10000005, "token time out")
	UnauthorizedTokenGenerate=NewError(10000006, "token generate error")
	TooManyRequests         = NewError(10000007, "too many request")
)