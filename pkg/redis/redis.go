package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
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

func (r *Redis) Set(key string, value interface{}, ex time.Duration) error {
	return r.rdb.Set(context.Background(), key, value, ex).Err()
}

func (r *Redis) Get(key string) (string, error) {
	return r.rdb.Get(context.Background(), key).Result()
}

func (r *Redis) HExists(key string, field string) (bool, error) {
	return r.rdb.HExists(context.Background(), key, field).Result()
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

func (r *Redis) Exists(key string) (bool, error) {
	result, err := r.rdb.Exists(context.Background(), key).Result()
	return result != 0, err
}

func (r *Redis) LPush(key string, value ...interface{}) error {
	return r.rdb.LPush(context.Background(), key, value...).Err()
}

func (r *Redis) LRem(key string, value interface{}) error {
	return r.rdb.LRem(context.Background(), key, 0, value).Err()
}

func (r *Redis) LRange(key string, start, stop int64) ([]string, error) {
	return r.rdb.LRange(context.Background(), key, start, stop).Result()
}

func (r *Redis) SAdd(key string, value ...interface{}) error {
	return r.rdb.SAdd(context.Background(), key, value...).Err()
}

func (r *Redis) SRem(key string, value ...interface{}) error {
	return r.rdb.SRem(context.Background(), key, value...).Err()
}

func (r *Redis) SMembers(key string) ([]string, error) {
	return r.rdb.SMembers(context.Background(), key).Result()
}

func (r *Redis) SIsMember(key string, value interface{}) (bool, error) {
	return r.rdb.SIsMember(context.Background(), key, value).Result()
}
