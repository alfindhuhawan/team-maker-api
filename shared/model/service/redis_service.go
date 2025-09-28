package service

import "context"

type RedisService interface {
	SetDataOnRedis(ctx context.Context, obj []byte, key string) error
	GetDataOnRedis(ctx context.Context, key string) (string, error)
}
