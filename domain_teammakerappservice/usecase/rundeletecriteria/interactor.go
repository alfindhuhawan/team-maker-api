package rundeletecriteria

import (
	"context"
)

//go:generate mockery --name Outport -output mocks/

type runDeleteCriteriaInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runDeleteCriteriaInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runDeleteCriteriaInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	//!

	// validation : first check the player exist or not
	_, err := r.outport.FindOneCriteria(ctx, req.CriteriaID)
	if err != nil {
		return nil, err
	}

	err = r.outport.DeleteCriteria(ctx, req.CriteriaID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
