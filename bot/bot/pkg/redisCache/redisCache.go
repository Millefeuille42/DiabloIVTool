package redisCache

import "github.com/redis/go-redis/v9"

func New(opts *redis.Options) *redis.Client {
	return redis.NewClient(opts)
}

var Client *redis.Client
