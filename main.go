package main

import (
	"fmt"
	"github.com/xinliangnote/go-fsm-order/internal/order"
)

func main() {
	// 通过订单ID 或 其他信息查询到订单状态
	orderStatus := order.StatusDefault

	orderMachine, err := order.NewFSM(orderStatus)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 创建订单，订单创建成功后再给用户发送短信
	if _, err = orderMachine.Call(order.EventCreate,
		order.WithOrderId(1),
		order.WithOrderName("测试订单"),
		order.WithHandlerSendSMS(sendSMS),
	); err != nil {
		fmt.Println(err.Error())
	}

	// 修改订单
	if _, err = orderMachine.Call(order.EventModify); err != nil {
		fmt.Println(err.Error())
	}

	// 确认订单
	if _, err = orderMachine.Call(order.EventConfirm); err != nil {
		fmt.Println(err.Error())
	}

	// 修改订单
	if _, err = orderMachine.Call(order.EventModify); err != nil {
		fmt.Println(err.Error())
	}

	// 确认订单
	if _, err = orderMachine.Call(order.EventConfirm); err != nil {
		fmt.Println(err.Error())
	}

	// 支付订单
	if _, err = orderMachine.Call(order.EventPay); err != nil {
		fmt.Println(err.Error())
	}
}

// sendSMS 发送短信，为了演示写在了这里
func sendSMS(mobile, content string) error {
	fmt.Println(fmt.Sprintf("发送短信，给(%s)发送了(%s)", mobile, content))
	return nil
}
