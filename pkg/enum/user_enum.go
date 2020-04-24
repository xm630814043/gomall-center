package enum

type UserType int

const (
	USER_TYPE_NOMAL   UserType = iota // 一般会员
	USER_TYPE_COMPANY                 // 企业会员
)
