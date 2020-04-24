package enum

// 交易类型 求购/供应
type BusinessType int8

const (
	BUSINESS_BUY  BusinessType = iota + 1 //求购
	BUSINESS_SELL                         // 供应
)

type BusinessCheckType int8

const (
	BUSINESS_CHECK_BOTH_NOT = iota //双方都没确认
	BUSINESS_CHECK_JIA             // 发布方确认
	BUSINESS_CHECK_YI              // 参与方确认
	BUSINESS_CHECK_BOTH            // 双方确认
)
