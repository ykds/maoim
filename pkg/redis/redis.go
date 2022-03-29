package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	rdb *redis.Client
}

func New(c *Config) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Host + ":" + c.Port,
		Username: c.User,
		Password: c.Password,
		DB:       c.DB,
	})
	return &Redis{
		rdb: rdb,
	}
}

func (r *Redis) HSet(key string, field string, value interface{}) error {
	return r.rdb.HSet(context.Background(), key, field, value).Err()
}

func (r *Redis) HGet(key string, field string) (string, error) {
	return r.rdb.HGet(context.Background(), key, field).Result()
}

func (r *Redis) HDel(key, field string) error {
	return r.rdb.HDel(context.Background(), key, field).Err()
}

