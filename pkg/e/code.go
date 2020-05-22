package e

const (
	SUCCESS      = 200
	ERROR        = 500
	Time_ERROR   = 1291
	BAD_REQUEST  = 400
	UNAUTHORIZED = 401
	FORBIDDEN    = 403

	// 业务编码从10000开始
	INCOMPLETE = 10000
	// account
	ERROR_EXIST_USER       = 10001
	ERROR_USER_OR_PASSWORD = 10002
	ERROR_VALIDCODE        = 10003
	ERROR_NOMOBILE         = 10004
	ERROR_NOPWD            = 10005

	// product
	ERROR_ProductData     = 10006
	ERROR_ProductSkuData  = 10007
	ERROR_ProductAttrData = 10008
)
