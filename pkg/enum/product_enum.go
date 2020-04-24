package enum

type PublishStatus int

const (
	PRODUCT_DRAFT    PublishStatus = iota // 草稿
	PRODUCT_AUDITING                      // 审核中
	PRODUCT_TO_ON                         // 待上架
	PRODUCT_ON                            // 已上架
	PRODUCT_OFF                           // 下架
)
