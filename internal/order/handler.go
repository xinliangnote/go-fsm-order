package order

import "fmt"

var (
	// handlerCreate 创建订单
	handlerCreate = Handler(func(opt *Opt) (State, error) {
		message := fmt.Sprintf("正在处理创建订单逻辑，订单ID(%d), 订单名称(%s) ... 处理完毕！", opt.OrderId, opt.OrderName)
		fmt.Println(message)

		if opt.HandlerSendSMS != nil {
			_ = opt.HandlerSendSMS("18888888888", "恭喜你预定成功了！")
		}

		return StatusReserved, nil
	})

	// handlerConfirm 确认订单
	handlerConfirm = Handler(func(opt *Opt) (State, error) {
		return StatusConfirmed, nil
	})

	// handlerModify 修改订单
	handlerModify = Handler(func(opt *Opt) (State, error) {
		return StatusReserved, nil
	})

	// handlerPay 支付订单
	handlerPay = Handler(func(opt *Opt) (State, error) {
		return StatusLocked, nil
	})
)
