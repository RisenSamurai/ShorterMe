package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func RedisInit() (*redis.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{
		Addr:        "redis:6379",
		Password:    "",
		DB:          0,
		PoolSize:    100,
		PoolTimeout: 60 * time.Second,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("redis init fail" + err.Error())
		return nil, err
	}

	log.Println("redis init success")

	return rdb, nil

}
