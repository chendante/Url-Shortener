package base

import (
	"Url-Shortener/const"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"time"
)

var Db *gorm.DB
var MRedis redis.Conn
var P sarama.SyncProducer

func init() {
	var err error
	Db, err = gorm.Open("mysql", _const.DataSourceName)
	if err != nil {
		fmt.Print(err.Error())
	}

	MRedis, err = redis.Dial("tcp", _const.Address)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	config.Version = sarama.V0_10_0_0
	P, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
}