package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var offset	int64
var consumer sarama.Consumer

func main() {
	initConsumer()
	TestConsumer()
}

func initConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	var err error
	consumer, err = sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	offset = 3
}

func TestConsumer() {
	partitionConsumer, err := consumer.ConsumePartition("test", 0, offset)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
			offset = msg.Offset
			// model.UpdateUrlVisits(string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
}
