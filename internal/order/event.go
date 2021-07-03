package order

// 定义订单事件
const (
	EventCreate  = Event("创建订单")
	EventConfirm = Event("确定订单")
	EventModify  = Event("修改订单")
	EventPay     = Event("支付订单")
)

// 定义订单事件对应的处理方法
var eventHandler = map[Event]Handler{
	EventCreate:  handlerCreate,
	EventConfirm: handlerConfirm,
	EventModify:  handlerModify,
	EventPay:     handlerPay,
}
