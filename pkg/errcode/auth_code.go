package errcode

var (
	ErrorAuthExist = NewError(30000001, "user has exist")
	ErrorAuthCreateFail = NewError(30000002, "create user error")
)
