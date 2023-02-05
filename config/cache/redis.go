package cache

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	rdb *redis.Client
}

type Redis interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
}

func InitRedis() Redis {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))

	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)
	return &redisClient{
		rdb: rdb,
	}
}

func (c *redisClient) Set(ctx context.Context, key string, value interface{}) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = c.rdb.Set(ctx, key, jsonValue, 10 * time.Minute).Err()
	return err
}

func (c *redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}