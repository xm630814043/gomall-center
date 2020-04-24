package enum

type CompanyStatus int8

const (
	COMPANY_NO_THROUGH      CompanyStatus = iota - 1 // 审核不通过
	COMPANY_UNDER_REVIEW                             // 等待审核
	COMPANY_APPROVAL_GANTED                          // 审核通过
)

type CompanyType int8

const (
	COMPAY_NO_BUSINESS CompanyType = iota + 1
	COMPAY_YES_BUSINESS
)
