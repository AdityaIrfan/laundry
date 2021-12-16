package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func Connect() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
