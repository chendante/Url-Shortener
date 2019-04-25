package base

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

var Db *sql.DB

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
	Db, err = sql.Open("mysql", dataLoaded["dataSourceName"])
	if err != nil {
		fmt.Print(err.Error())
	}
	err = Db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
}