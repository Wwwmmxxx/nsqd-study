package main

func main() {

	var (
		topic      = "nsq"
		channel    = "testGo"
		nsqAddr    = "localhost:4150"
		lookupAddr = "localhost:4161"
	)

	NewProducer(nsqAddr, topic)
	NewConsumer(lookupAddr, nsqAddr, topic, channel)

}
