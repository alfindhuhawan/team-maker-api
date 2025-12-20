package runcreatecriteria

import (
	"context"
	"team-maker-api/shared/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type runCreateCriteriaInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runCreateCriteriaInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runCreateCriteriaInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	// save criteria
	err = r.outport.SaveCriteria(ctx, &entity.Criteria{
		Title:        req.Title,
		TitleAliases: req.TitleAliases,
		// Rank:        req.Rank,
		// Rule:        req.Rule,
		Description: req.Description,
		CreatedAt:   req.Now,
		CreatedBy:   "",
		UpdatedAt:   req.Now,
		UpdatedBy:   "",
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
