package enum

type CategoryStatus int
type CategoryType int

const (
	// 上线
	CATEGORY_OFF_LINE CategoryStatus = iota
	// 下线
	CATEGORY_ON_LINE
)

const (
	CATEGORY_ALL CategoryType = iota
	CATEGORY_DEFAULT
)
