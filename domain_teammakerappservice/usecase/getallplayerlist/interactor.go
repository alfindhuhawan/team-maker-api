package getallplayerlist

import (
	"context"
)

//go:generate mockery --name Outport -output mocks/

type getAllPlayerListInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllPlayerListInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllPlayerListInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	dataPlayers, err := r.outport.FindAllPlayerList(ctx, nil)
	if err != nil {
		return nil, err
	}

	res.Items = dataPlayers

	return res, nil
}
