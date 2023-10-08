package publish

import (
	"2102Aweek3/inits"
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func Publish(topic string, body string) bool {
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{inits.RocketServer}))
	if err != nil {
		panic(" producer ")
	}
	p.Start()
	defer p.Shutdown()

	_, err = p.SendSync(context.Background(), primitive.NewMessage(topic, []byte(body)))
	if err != nil {
		fmt.Printf(" : %s\n", err)
		return false
	}
	return true
}
