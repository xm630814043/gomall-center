package enum

type Order_Staus int

const (
	OrderCreate   Order_Staus = iota
	OrderNoPay                // 待付款
	OrderNoSend               // 待发货
	OrderNoAccept             // 待收货
	OrderNoCommon             // 待评价
	OrderFinish               // 完成
	OrderCancel               // 取消
)
