package main

import (
	"2102Amonth/common/database"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

func main() {
	Consumer()
}

func Consumer() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("monthexam"),
	)

	if err := c.Subscribe("refund_order", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("获取到值： %v \n", msgs[i])

			// 消费信息
			m := msgs[i].Message.Body
			fmt.Printf("获取到值： %v \n", m)

			var mm map[string]interface{}
			json.Unmarshal(m, &mm)
			orderNo := mm["OrderNo"]
			orderInfo := database.GetOrderInfo(orderNo.(string))
			orderInfo.Status = 3

			database.UpdateOrder(orderInfo)

		}
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		fmt.Println("读取消息失败")
	}
	_ = c.Start()
	//不能让主goroutine退出
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}
