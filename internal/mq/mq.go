package mq

import (
	"chatroom/internal/common"
	"chatroom/internal/config"
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

var (
	G_Admin        admin.Admin
	G_Producer     rocketmq.Producer
	G_PushConsumer rocketmq.PushConsumer

	topics = [...]string{common.MessageTopic, common.ReadTopic}
)

func InitMQ() (err error) {
	rlog.SetLogLevel("error")
	fmt.Println(config.G_Config.Rocketmq.Endpoints)
	G_Admin, err = admin.NewAdmin(
		admin.WithResolver(primitive.NewPassthroughResolver(config.G_Config.Rocketmq.Endpoints)),
	)
	if err != nil {
		return err
	}
	fmt.Println("创建topic：", topics)
	for _, topic := range topics {
		if err = G_Admin.CreateTopic(
			context.Background(),
			admin.WithTopicCreate(topic),
			admin.WithBrokerAddrCreate(config.G_Config.Rocketmq.BrokerAddr),
		); err != nil {
			fmt.Println("err", err)
			return err
		}
	}

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
