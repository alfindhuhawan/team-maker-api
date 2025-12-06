package rundeleteplayer

import (
	"context"
	"team-maker-api/shared/model/enum"
	"team-maker-api/shared/model/repository"
)

//go:generate mockery --name Outport -output mocks/

type runDeletePlayerInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runDeletePlayerInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runDeletePlayerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// validation : first check the player exist or not
	_, err := r.outport.FindOnePlayer(ctx, enum.IDFilterByEnum, repository.FilterPlayer{
		ID: req.PlayerID,
	})
	if err != nil {
		return nil, err
	}

	err = r.outport.DeletePlayer(ctx, req.PlayerID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
