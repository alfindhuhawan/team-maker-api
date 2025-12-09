package runupdateplayer

import (
	"context"
	"fmt"
	"team-maker-api/shared/model/enum"
	"team-maker-api/shared/model/repository"
)

//go:generate mockery --name Outport -output mocks/

type runUpdatePlayerInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runUpdatePlayerInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runUpdatePlayerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	playerDataOld, err := r.outport.FindOnePlayer(ctx, enum.IDFilterByEnum, repository.FilterPlayer{
		ID: req.PlayerID,
	})
	if err != nil {
		return nil, err
	}

	// must check its from the same id or not
	if playerDataOld.PlayerCode != req.PlayerCode {
		playerExist, _ := r.outport.FindOnePlayer(ctx, enum.PlayerCodeFilterByEnum, repository.FilterPlayer{
			PlayerCode: req.PlayerCode,
		})
		if playerExist != nil {
			return nil, fmt.Errorf("player code has been used")
		}
	}

	err = r.outport.UpdatePlayer(ctx, req.PlayerID, &repository.UpdatePlayerRequest{
		Name:       req.Name,
		PlayerRank: req.PlayerRank,
		PlayerCode: req.PlayerCode,
		Criteria:   req.Criteria,
		UpdatedAt:  req.TimeNow,
		UpdatedBy:  "-", // TODO : must change when auth have been created
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
