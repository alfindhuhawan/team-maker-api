package runcreateteam

import "team-maker-api/shared/model/repository"

// Outport of usecase
type Outport interface {
	repository.FindAllPlayerListRepo
	repository.CreateBalanceMatchmakingTeamRepo
}
