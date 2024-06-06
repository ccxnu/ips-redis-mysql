package database

import (
	"context"

	"github.com/ccxnu/ips-redis-mysql/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisConnection(c config.Config) (*redis.Client, error) {
	opt, err := redis.ParseURL(c.Redis.Url)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
