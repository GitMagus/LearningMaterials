package main

import (
	"2102Amonth/common/database"
	"2102Amonth/common/rocketmq"
	"encoding/json"
)

func main() {
	// 遍历超时订单，压入消息队列
	PushToRocketmq()
}

func PushToRocketmq() {
	// 获取超时订单
	refundOrderList := database.GetRefundOrderList()
	for _, val := range refundOrderList {
		b, _ := json.Marshal(val)
		rocketmq.PublishMsg(b, "refund_order")
	}
}
