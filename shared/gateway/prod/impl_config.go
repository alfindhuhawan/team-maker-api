package prod

import (
	"context"
	"team-maker-api/shared/infrastructure/config"
	"team-maker-api/shared/infrastructure/logger"
)

type ConfigImpl struct {
	Log logger.Logger
	Cfg *config.Config
}

func (r *ConfigImpl) GetLoginConfig(ctx context.Context) (string, string) {
	keyConfig := r.Cfg.AuthKey.KeyConfig
	jwtSecret := r.Cfg.AuthKey.JWTSecret

	return keyConfig, jwtSecret
}

func (r *ConfigImpl) GetConfigExceptionServiceType(ctx context.Context) []string {
	exceptionServiceType := r.Cfg.ExceptionServiceType

	return exceptionServiceType
}
