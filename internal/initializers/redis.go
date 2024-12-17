package initializers

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

var ctx = context.TODO()

func ConnectToRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: CONFIG.REDIS_HOST + ":" + CONFIG.REDIS_PORT,
		// Username: CONFIG.REDIS_USER,
		Password: CONFIG.REDIS_PASSWORD,
		DB:       0,
		PoolSize: 1000,
	})

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		fmt.Printf("Redis connection Error:\n %v", err)
	} else {
		fmt.Println("Connected to redis!")
	}
}

