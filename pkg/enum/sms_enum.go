package enum

// SMSType 短信的类型
type SMSType int

const (
	_        SMSType = iota
	SMSLOGIN         // 短信登录
	SMSAUTH          // 身份认证，比如注册
	SMSPAY           // 支付认证，预留
)
