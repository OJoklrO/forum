package errcode

var (
	ErrorCountPostsFail = NewError(20010001, "count posts error")
	ErrorGetPostFail = NewError(20010002, "get tag list error")
	ErrorGetPostListFail = NewError(20010003, "count tag error")
	ErrorCreatePostFail = NewError(20010004, "create tag error")
	ErrorUpdatePostFail = NewError(20010005, "update tag error")
	ErrorDeletePostFail = NewError(20010006, "delete tag error")

)
