package base

import (
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
	bytes, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
	}
	var dataLoaded map[string]string
	if err := json.Unmarshal(bytes, &dataLoaded); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
	}
	Db, err = gorm.Open("mysql", dataLoaded["dataSourceName"])
	if err != nil {
		fmt.Print(err.Error())
	}

	MRedis, err = redis.Dial("tcp", dataLoaded["address"])
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	P, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
}