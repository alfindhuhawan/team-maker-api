package repository

import (
	"context"
	"team-maker-api/shared/model/entity"
)

type CreateBalanceMatchmakingTeamRepo interface {
	CreateBalanceMatchmakingTeam(ctx context.Context, totalTeam int, maxPlayer int, dataPlayers []*entity.Player) ([][]string, error)
}
