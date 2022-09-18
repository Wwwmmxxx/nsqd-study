package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
)

func NewConsumer(lookupAddr, nsqAddr, topic, channel string) {

	// 定义nsq生产者配置
	config := nsq.NewConfig()

	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Println("NewConsumer error", err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Print(string(message.Body))
		return nil
	}))

	// 连接到nsq lookup
	//err = consumer.ConnectToNSQLookupd(lookupAddr)
	//if err != nil {
	//	fmt.Println("ConnectToNSQlLookup error", err)
	//}

	// 连接到nsq
	err = consumer.ConnectToNSQD(nsqAddr)
	if err != nil {
		fmt.Println("ConnectToNSQD error", err)
	}

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt

	// 停止链接
	consumer.Stop()
}
