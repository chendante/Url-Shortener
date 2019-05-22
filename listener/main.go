package main

import (
	"Url-Shortener/model"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/robfig/cron"
	"time"
)

var offset	int64
var consumer sarama.Consumer

func main() {
	c := cron.New()
	initConsumer()
	err := c.AddFunc("0 0 2 * * ?", TestConsumer)	// 凌晨两点执行一次
	if err != nil {
		fmt.Printf("c.AddFunc error %s\n", err.Error())
	}
	c.Start()
	select {}
}

func initConsumer() {
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_0_0
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	var err error
	consumer, err = sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	offset = sarama.OffsetOldest
	fmt.Print("offset: ", offset)
}

func TestConsumer() {
	partitionConsumer, err := consumer.ConsumePartition("test", 0, offset)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	now := time.Now()
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
			if now.Unix() > msg.Timestamp.Unix() {
				offset = msg.Offset
			} else {
				break ConsumerLoop
			}
			model.UpdateUrlVisits(string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		case <-time.After(time.Second * 10):
			// 超时退出
			break ConsumerLoop
		}
	}
	offset += 1
	partitionConsumer.Close()
}
