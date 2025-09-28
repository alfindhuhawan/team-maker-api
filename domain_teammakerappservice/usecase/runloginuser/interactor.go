package runloginuser

import (
	"context"
)

//go:generate mockery --name Outport -output mocks/

type runLoginUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runLoginUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runLoginUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	// code your usecase definition here ...
	//!

	return res, nil
}
