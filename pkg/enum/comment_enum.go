package enum

type CommentStatus int8

const (
	COMMENT_NOT_PASS  CommentStatus = iota - 1 //不通过
	COMMENT_WAIT_PASS                          // 等待审核
	COMMENT_PASS                               // 通过
)
