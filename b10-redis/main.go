package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	set, err := rdb.SetNX(ctx, "key", "value", 20*time.Second).Result()

	fmt.Println(set, err)

	val, err := rdb.Get(ctx, "key").Result()

	fmt.Println(val, err)

	del, err := rdb.Del(ctx, "key").Result()

	fmt.Println(del, err)
}
