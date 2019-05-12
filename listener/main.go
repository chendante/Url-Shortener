package main

import (
	"Url-Shortener/model"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	TestConsumer()
}

func TestConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
			model.UpdateUrlVisits(string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
}
