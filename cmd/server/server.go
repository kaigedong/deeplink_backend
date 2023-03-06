package main

import (
	"context"
	"deeplink_backend/pkg/router"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "myPassword",
		DB:       0, // use default DB
	})

	var ctx = context.Background()
	if _, err = rdb.Ping(ctx).Result(); err != nil {
		return err
	}

	if err := rdb.Set(ctx, "key", "value", 0).Err(); err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	return nil
}

func init() {}

func main() {
	viper.SetDefault("port", ":8000")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Using defualt config...")
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

	}

	engine := router.Default()
	engine.Run(viper.GetString("port"))
}
