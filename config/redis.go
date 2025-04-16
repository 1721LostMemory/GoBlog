package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RedisDB *redis.Client

func InitRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.address"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})

	ctx := context.Background()
	pong, err := RedisDB.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("redis连接失败", pong, err))
	} else {
		fmt.Println("redis连接成功", pong)
	}
}
