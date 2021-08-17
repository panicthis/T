package T

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func Visit() {
	fmt.Println("T 1.0.0")
}

func UseRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}
