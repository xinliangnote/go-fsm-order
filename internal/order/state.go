package order

// 定义订单状态
const (
	StatusDefault   = State(0)
	StatusReserved  = State(10)
	StatusConfirmed = State(20)
	StatusLocked    = State(30)
)

// statusText 定义订单状态文案
var statusText = map[State]string{
	StatusDefault:   "默认",
	StatusReserved:  "已预订",
	StatusConfirmed: "已确认",
	StatusLocked:    "已锁定",
}

// statusEvent 定义订单状态对应的可操作事件
var statusEvent = map[State][]Event{
	StatusDefault:   {EventCreate},
	StatusReserved:  {EventConfirm},
	StatusConfirmed: {EventModify, EventPay},
}

func StatusText(status State) string {
	return statusText[status]
}
