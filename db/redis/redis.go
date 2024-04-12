package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis"
)

func init() {
	InitRedis()
}

var (
	Client *redis.Client
	Nil    = redis.Nil
)

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 访问密码，如果没有则为空字符串
		DB:       0,                // 使用的 Redis 数据库编号
	})

	_, err := Client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func JWTMatch(ctx context.Context, username, token string) bool {
	jwt, err := Client.WithContext(ctx).Get(username).Result()
	if err == Nil || jwt != token {
		return false
	}
	return true
}
