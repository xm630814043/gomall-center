package enum

type IsFreeFeeType int8

const (
	T_ALOE_GEL IsFreeFeeType = iota
	T_Pinkage
)

type HasFreeConditionType int8

const (
	T_NO_CONDITIONS HasFreeConditionType = iota
	T_YES_CONDITIONS
)
