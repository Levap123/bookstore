package redis

import (
	"github.com/Levap123/book_service/internal/configs"
	"github.com/go-redis/redis/v8"
)

func InitRedis(cfg *configs.Configs) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})
}
