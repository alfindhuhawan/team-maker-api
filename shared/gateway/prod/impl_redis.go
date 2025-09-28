package prod

import (
	"context"
	"team-maker-api/shared/infrastructure/cache"
	"team-maker-api/shared/infrastructure/config"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/model/errorenum"
	"time"
)

type RedisImpl struct {
	Log         logger.Logger
	Cfg         *config.Config
	CacheClient cache.Cache
}

func (r *RedisImpl) SetDataOnRedis(ctx context.Context, obj []byte, key string) error {

	// SET DATA ON REDIS
	err := r.CacheClient.Set(ctx, key, obj, time.Duration(r.Cfg.Cache.Redis.TimeDuration)*time.Minute) // TODO : MASUKIN CONFIG 1 JAMNYA
	if err != nil {
		return errorenum.RedisErrorEN.Var("SET METHOD")
	}

	return nil
}

func (r *RedisImpl) GetDataOnRedis(ctx context.Context, key string) (string, error) {
	value, err := r.CacheClient.Get(ctx, key)
	if err != nil {
		return "", errorenum.RedisErrorEN.Var("GET METHOD")
	}

	return value, nil
}

// func SetDataSettingsOnRedis(ctx context.Context, cacheClient cache.Cache, cfg *config.Config, obj *service.ShipmentSettingsResponse, key string) error {
// 	byteObj, err := json.Marshal(obj)
// 	if err != nil {
// 		return errorenum.RedisErrorEN.Var("SET METHOD")
// 	}

// 	// SET DATA ON REDIS
// 	err = cacheClient.Set(ctx, key, byteObj, time.Duration(cfg.Cache.Redis.TimeDurationSetting)*time.Minute) // TODO : MASUKIN CONFIG 1 JAMNYA
// 	if err != nil {
// 		return errorenum.RedisErrorEN.Var("SET METHOD")
// 	}

// 	return nil
// }

// func (r *RedisImpl) PurgeRedis(ctx context.Context) error {
// 	// SET DATA ON REDIS
// 	err := r.CacheClient.Purge(ctx) // TODO : MASUKIN CONFIG 1 JAMNYA
// 	if err != nil {
// 		fmt.Println("response purge here >>")
// 		fmt.Println(err)
// 		return err
// 	}

// 	return nil
// }
