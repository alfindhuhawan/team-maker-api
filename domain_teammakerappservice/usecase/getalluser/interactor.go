package getalluser

import (
	"context"
	"team-maker-api/shared/model/repository"
)

//go:generate mockery --name Outport -output mocks/

type getAllUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	dataUsers, count, err := r.outport.FindAllUser(ctx, repository.FindAllUserRequest{
		Page:     req.Page,
		Size:     req.Size,
		Name:     req.Name,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = dataUsers

	return res, nil
}
