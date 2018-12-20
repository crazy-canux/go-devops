package main

import (
    "github.com/go-redis/redis"
    "fmt"
)


func main() {
    client := redis.NewClient(
        &redis.Options{
            Addr: "127.0.0.1:6379",
            Password: "",
            DB: 0,
        },
    )
    defer client.Close()

    //pong, err := client.Ping().Result()
    //fmt.Println(pong, err)

    info, err := client.Info().Result()
    fmt.Println(info, err)
}