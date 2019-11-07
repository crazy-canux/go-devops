package nosql

import (
    "github.com/go-redis/redis"
    "log"
)

func Info(ip, pw string, db int) (string, error) {
    client := redis.NewClient(
        &redis.Options{
            Addr: ip,
            Password: pw,
            DB: db,
        },
    )
    defer client.Close()

    ping, err := client.Ping().Result()
    if err != nil {
        log.Printf("Connection failed: %s", ping)
        return "", err
    }
    info, err := client.Info().Result()
    if err != nil {
        log.Printf("Get info failed: %s", info)
        return "", err
    }
    return info, nil
}