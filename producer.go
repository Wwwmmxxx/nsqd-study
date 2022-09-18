package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

func NewProducer(nsqAddr, topic string) {

	// 定义nsq生产者配置
	config := nsq.NewConfig()

	// 创建生产者
	producer, err := nsq.NewProducer(nsqAddr, config)
	if err != nil {
		fmt.Println("New Producer error", err)
	}

	for i := 0; i < 10; i++ {
		// 发送消息
		producer.Publish(topic, []byte(fmt.Sprintf("push message %d\n", i)))
	}

	// 停止链接
	producer.Stop()
}
