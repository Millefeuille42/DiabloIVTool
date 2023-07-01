package redisCache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func NewContext() context.Context {
	return context.Background()
}

func New(opts *redis.Options) *redis.Client {
	return redis.NewClient(opts)
}

var Context context.Context
var Client *redis.Client
