package errcode

var (
	ErrorCountPostsFail = NewError(20010001, "count posts error")
	ErrorGetPostFail = NewError(20010002, "get post error")
	ErrorGetPostListFail = NewError(20010003, "get post list error")
	ErrorCreatePostFail = NewError(20010004, "create post error")
	ErrorUpdatePostFail = NewError(20010005, "update post error")
	ErrorDeletePostFail = NewError(20010006, "delete post error")

)
