package mq

import (
	"chatroom/internal/config"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

var (
	G_Producer     rocketmq.Producer
	G_PushConsumer rocketmq.PushConsumer
)

func InitMQ() (err error) {
	rlog.SetLogLevel("error")
	G_Producer, err = rocketmq.NewProducer(
		producer.WithNameServer(config.G_Config.Rocketmq.Endpoints),
		producer.WithRetry(2),
		producer.WithGroupName("ProducerGroup"),
		//根据msg设置的ShardingKey来选择写入到的queue
		producer.WithQueueSelector(producer.NewHashQueueSelector()),
	)
	if err != nil {
		return err
	}
	G_PushConsumer, err = rocketmq.NewPushConsumer(
		consumer.WithNameServer(config.G_Config.Rocketmq.Endpoints),
		consumer.WithGroupName("PushConsumerGroup"),
	)
	if err != nil {
		return err
	}
	return err
}
