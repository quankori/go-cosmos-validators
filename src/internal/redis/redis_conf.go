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

func GetMessage(key string) []string {
	redisClient := redisClient()
	result, _ := redisClient.BLPop(0*time.Second, key).Result()
	fmt.Println(result)

	return result
}
