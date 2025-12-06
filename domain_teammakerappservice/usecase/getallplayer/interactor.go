package getallplayer

import (
	"context"
	"team-maker-api/shared/model/repository"
)

//go:generate mockery --name Outport -output mocks/

type getAllPlayerInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllPlayerInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllPlayerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	dataPlayers, count, err := r.outport.FindAllPlayer(ctx, repository.FindAllPlayerRequest{
		Page:       req.Page,
		Size:       req.Size,
		Name:       req.Name,
		PlayerRank: req.PlayerRank,
	})
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = dataPlayers

	return res, nil
}
