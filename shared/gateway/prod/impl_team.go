package prod

import (
	"context"
	"team-maker-api/shared/infrastructure/database"
	"team-maker-api/shared/infrastructure/logger"
	"team-maker-api/shared/model/entity"
)

type TeamImpl struct {
	*database.MongoWithTransaction
	Log    logger.Logger
	DbName string
}

func (r *TeamImpl) CreateBalanceMatchmakingTeam(ctx context.Context, totalTeam int, maxPlayer int, dataPlayers []*entity.Player) ([][]string, error) {
	var dataTeam [][]string

	// lenDataPlayers := len(dataPlayers)

	// initTotalTeam := 0
	// initLowRankPlayer := make(map[int]bool, 0)
	// for lenDataPlayers > 0 {
	// 	randomIndex := util.GenerateRandomNumber(lenDataPlayers)

	// 	player := dataPlayers[randomIndex]

	// 	if  {

	// 	}

	// 	// decrease len
	// 	lenDataPlayers -= 1
	// }

	return dataTeam, nil
}
