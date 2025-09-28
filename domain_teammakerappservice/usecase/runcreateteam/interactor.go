package runcreateteam

import (
	"context"
	"team-maker-api/shared/model/enum"
	"team-maker-api/shared/model/repository"
)

//go:generate mockery --name Outport -output mocks/

type runCreateTeamInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runCreateTeamInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runCreateTeamInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	dataPlayers, err := r.outport.FindAllPlayerList(ctx, &repository.FindAllPlayerListRequest{
		PlayerIDs: req.PlayerIDs,
	})
	if err != nil {
		return nil, err
	}

	var dataTeams [][]string
	if req.MatchmakingType == enum.BalanceMatchmakingTypeEnum {
		dataTeams, err = r.outport.CreateBalanceMatchmakingTeam(ctx, req.TotalTeam, req.MaxPlayer, dataPlayers)
		if err != nil {
			return nil, err
		}
	}

	res.Items = dataTeams

	return res, nil
}
