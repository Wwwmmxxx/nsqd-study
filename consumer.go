package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

func NewConsumer(lookupAddr, nsqAddr, topic, channel string) {

	// 定义nsq生产者配置
	config := nsq.NewConfig()

	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Println("NewConsumer error", err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Println(string(message.Body))
		return nil
	}))

	// 连接到nsq lookup
	//err = consumer.ConnectToNSQLookupd(lookupAddr)
	// 连接到nsq
	err = consumer.ConnectToNSQD(nsqAddr)
	if err != nil {
		fmt.Println("ConnectToNSQlLookup error", err)
	}

	//interrupt := make(chan os.Signal)
	//signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	//<-interrupt

	// 停止链接
	consumer.Stop()
}
