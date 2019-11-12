package nosql

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Status struct {
	Host            string                 `bson:"host"`
	Version         string                 `bson:"version"`
	Process         string                 `bson:"process"`
	Pid             int64                  `bson:"pid"`
	Uptime          int64                  `bson:"uptime"`
	UptimeMillis    int64                  `bson:"uptimeMillis"`
	UptimeEstimate  int64                  `bson:"uptimeEstimate"`
	Asserts         map[string]int64       `bson:"asserts"`
	ShardCursorType map[string]interface{} `bson:"shardCursorType"`
	StorageEngine   map[string]string      `bson:"storageEngine"`
}

func ServerStatus(ip string) (*Status, error) {
	var status = &Status{}
	uri := fmt.Sprintf("mongodb://%s", ip)
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Println("Dial failed.")
		return status, err
		fmt.Println("dial failed: ", err)
	}
	defer session.Close()

	err = session.Ping()
	if err != nil {
		log.Println("ping failed.")
		return status, err
	}

	err = session.Run(
		bson.D{
			{Name: "serverStatus"},
			{Value: 1},
		},
		status,
	)
	if err != nil {
		log.Println("Get serverStatus failed.")
		return status, err
	}
	return status, err
}
