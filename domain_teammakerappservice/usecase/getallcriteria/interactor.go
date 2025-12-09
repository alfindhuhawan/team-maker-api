package getallcriteria

import (
	"context"
	"team-maker-api/shared/model/repository"
)

//go:generate mockery --name Outport -output mocks/

type getAllCriteriaInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllCriteriaInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllCriteriaInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	dataCriterias, count, err := r.outport.FindAllCriteria(ctx, repository.FindAllCriteriaRequest{
		Page:  req.Page,
		Size:  req.Size,
		Title: req.Title,
	})
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = dataCriterias

	return res, nil
}
