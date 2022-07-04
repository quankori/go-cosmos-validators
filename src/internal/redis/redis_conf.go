package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0,
	})
}

func PushMessage(key string, value string) error {
	redisClient := redisClient()

	err := redisClient.RPush(key, value).Err()

	if err != nil {
		return err
	}
	return nil
}

func GetMessage(key string) chan string {
	r := make(chan string)
	go func() {
		for {
			redisClient := redisClient()
			result, err := redisClient.BLPop(0*time.Second, key).Result()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				r <- result[1]
			}
		}
	}()
	return r
}
