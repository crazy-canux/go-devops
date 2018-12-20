package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type ServerStatus struct {
	Host               string                 `bson:"host"`
	Version            string                 `bson:"version"`
	Process            string                 `bson:"process"`
	Pid                int64                  `bson:"pid"`
	Uptime             int64                  `bson:"uptime"`
	UptimeMillis       int64                  `bson:"uptimeMillis"`
	UptimeEstimate     int64                  `bson:"uptimeEstimate"`
	Asserts            map[string]int64       `bson:"asserts"`
	ShardCursorType    map[string]interface{} `bson:"shardCursorType"`
	StorageEngine      map[string]string      `bson:"storageEngine"`
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		fmt.Println("dial failed: ", err)
	}
	defer session.Close()

	//err = session.Ping()
	//if err != nil {
	//	fmt.Println("ping failed: ", err)
	//}

	//names, err := session.DatabaseNames()
	//if err != nil {
	//	fmt.Println("get databasenames failed: ", err)
	//}
	//fmt.Println(names)
	var serverResult = &ServerStatus{}
	session.Run(
		bson.D{
			{Name: "serverStatus"},
			{Value: 1},
		},
		serverResult)
	fmt.Println(*serverResult)
}