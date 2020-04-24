package enum

type PayMethod int

const (
	OFFLINE     PayMethod = iota // 线下支付
	ONLINE_BANK                  // 银联转账
	WEIXIN                       // 微信
	ZHIFUBAO                     // 支付宝
)
