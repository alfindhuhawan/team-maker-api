package repository

import (
	"context"
)

type GetLoginConfigRepo interface {
	GetLoginConfig(ctx context.Context) (string, string)
}

type GetConfigExceptionServiceTypeRepo interface {
	GetConfigExceptionServiceType(ctx context.Context) []string
}
