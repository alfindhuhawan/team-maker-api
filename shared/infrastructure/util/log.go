package util

import (
	"context"
	"team-maker-api/shared/infrastructure/logger"
)

func InsertLog(ctx context.Context, log logger.Logger, err error) {
	log.Error(ctx, err.Error())
}
