package cache

import (
	"context"
	"errors"
	"fmt"
	"team-maker-api/shared/infrastructure/config"
	"time"

	"github.com/go-redis/redis/v8"
)

// Cache We only use 3 base function in redis
type Cache interface {

	// Set put the initial value
	Set(ctx context.Context, key string, value []byte, exp time.Duration) error

	// Get the value
	Get(ctx context.Context, key string) (string, error)

	// Del Delete the value
	Del(ctx context.Context, key string) error

	// Exist Check existing key
	Exist(ctx context.Context, key string) (bool, error)

	// Exist Check existing key
	Purge(ctx context.Context) error
}

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCacheDefault(cfg *config.Config) Cache {

	redisOptions := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Cache.Redis.Address, cfg.Cache.Redis.Port),
		Password: cfg.Cache.Redis.Password,
		DB:       cfg.Cache.Redis.Db,
	}

	//if cfg.Environment == "PROD" {
	//	certs := x509.NewCertPool()
	//
	//	pemData, err := ioutil.ReadFile("certificate-cert.pem")
	//	if err != nil {
	//		panic(fmt.Errorf("%s", err))
	//	}
	//
	//	certs.AppendCertsFromPEM(pemData)
	//
	//	redisOptions.TLSConfig = &tls.Config{
	//		RootCAs: certs,
	//	}
	//}

	return &RedisCache{
		Client: redis.NewClient(&redisOptions),
	}
}

// Set receive key and value as input and return error
func (c *RedisCache) Set(ctx context.Context, key string, value []byte, exp time.Duration) error {
	err := c.Client.Set(ctx, key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get receive key and bytes (return object from passing by reference)
// return isExist flag and error
func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := c.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("Redis key %s not found", key)
	}

	if err != nil {
		return "", err
	}
	return result, nil
}

// Exist check the existence of key
// return isExist flag and error
func (c *RedisCache) Exist(ctx context.Context, key string) (bool, error) {

	_, err := c.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// Del deletes by key
func (c *RedisCache) Del(ctx context.Context, key string) error {

	err := c.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil

}

// Reset receive key and value as input and return error
func (c *RedisCache) Reset(ctx context.Context, key string, value []byte) error {
	err := c.Client.Set(ctx, key, value, redis.KeepTTL).Err()
	if err != nil {
		return err
	}
	return nil
}

// PURGE
func (c *RedisCache) Purge(ctx context.Context) error {
	err := c.Client.FlushAll(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}
