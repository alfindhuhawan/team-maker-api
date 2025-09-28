package prod

import (
	"team-maker-api/shared/infrastructure/config"
	"team-maker-api/shared/infrastructure/logger"
)

type ValidationImpl struct {
	Log logger.Logger
	Cfg *config.Config
}

func (r *ValidationImpl) CheckInArray(target string) bool {
	exceptionServiceType := r.Cfg.ExceptionServiceType

	for _, value := range exceptionServiceType {
		if value == target {
			return true
		}
	}

	return false
}
