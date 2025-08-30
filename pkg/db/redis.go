package db

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

//	&redis.Options{
//			Addr:     "localhost:6379",
//			Password: "",
//			DB:       0,
//		}

type RedisOptions = redis.Options

func NewRedis(opt *RedisOptions) (*Redis, error) {
	c := redis.NewClient(opt)
	_, err := c.Ping(context.Background()).Result()
	return &Redis{c}, err
}

func (r *Redis) Insert(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return r.client.Set(ctx, key, value, exp).Err()
}

func (r *Redis) Read(ctx context.Context, key string) (interface{}, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *Redis) Close() {
	r.client.Close()
}
