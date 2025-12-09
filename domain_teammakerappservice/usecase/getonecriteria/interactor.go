package getonecriteria

import (
	"context"
)

//go:generate mockery --name Outport -output mocks/

type getOneCriteriaInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getOneCriteriaInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getOneCriteriaInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	dataCriteria, err := r.outport.FindOneCriteria(ctx, req.CriteriaID)
	if err != nil {
		return nil, err
	}

	res.Item = dataCriteria

	return res, nil
}
