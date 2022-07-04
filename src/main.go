package main

import (
	"fmt"

	"github.com/quankori/go-cosmos-validators/src/internal/redis"
)

func main() {
	redis.PushMessage("key", "value 4")
	val := redis.GetMessage("key")
	redis.PushMessage("key", "value 123")
	fmt.Println(<-val)
}
