package mq

import (
	"chatroom/internal/config"
	"fmt"
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
	fmt.Println(config.G_Config.Rocketmq.Endpoints)
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
	G_Producer.Start()
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
	//G_PushConsumer.Subscribe("message", consumer.MessageSelector{})
	//G_PushConsumer.Start()
	//if err != nil {
	//	return err
	//}

	//global.MQProducer, err = rmq_client.NewProducer(&rmq_client.Config{
	//	Endpoint: config.G_Config.Rocketmq.Endpoint,
	//})
	//
	//global.MQConsumer, err = rmq_client.NewSimpleConsumer(&rmq_client.Config{
	//	Endpoint:      config.G_Config.Rocketmq.Endpoint,
	//	ConsumerGroup: "message",
	//})

	return err
}
