package errcode

var (
	ErrorCountPostsFail = NewError(20010001, "count posts error")
	ErrorGetPostFail = NewError(20010002, "get post error")
	ErrorGetPostListFail = NewError(20010003, "get post list error")
	ErrorCreatePostFail = NewError(20010004, "create post error")
	ErrorUpdatePostFail = NewError(20010005, "update post error")
	ErrorDeletePostFail = NewError(20010006, "delete post error")

	ErrorListCommentsFail = NewError(20020001, "list comments error")
	ErrorCountCommentsFail = NewError(20020002, "count comments error")
	ErrorCreateCommentFail = NewError(20020003, "create comment error")
	ErrorDeleteCommentFail = NewError(20020004, "delete comment error")
)
