package main

import (
	service "2102Aweek3/common"
	"2102Aweek3/inits"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"strconv"
	"time"
)

// 订阅消息处理
func main() {
	Consumer("score", "week3")
}

func Consumer(score string, group string) {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{inits.RocketServer}),
		consumer.WithGroupName(group),
	)

	if err := c.Subscribe(score, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("获取到值： %v \n", msgs[i])

			// 消费信息
			m := msgs[i].Message.Body
			fmt.Printf("获取到值： %v \n", m)

			var mm map[string]interface{}

			json.Unmarshal(m, &mm)
			userId := mm["user_id"]
			//oType := mm["type"]
			//integral := mm["integral"]
			// 调用随机数方法，生成新积分
			// 更新用户积分
			newScore := service.GetRandNum()
			uid, _ := strconv.Atoi(userId.(string))
			inits.UpdateUserIntegral(uid, newScore)
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
