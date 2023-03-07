package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "myPassword",
		DB:       0, // use default DB
	})

	var ctx = context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(fmt.Errorf("Connection to Redis server failed: %s \n", err))
	}

	// if err := rdb.Set(ctx, "key", "value", 0).Err(); err != nil {
	// 	panic(err)
	// }

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }

	// return nil
}
