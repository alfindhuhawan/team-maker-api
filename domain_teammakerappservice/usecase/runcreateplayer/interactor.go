package runcreateplayer

import (
	"context"
	"fmt"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/enum"
	"team-maker-api/shared/model/repository"
)

//go:generate mockery --name Outport -output mocks/

type runCreatePlayerInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runCreatePlayerInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runCreatePlayerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	// validation : check player code first, player code cant be duplicate
	playerExist, _ := r.outport.FindOnePlayer(ctx, enum.PlayerCodeFilterByEnum, repository.FilterPlayer{
		PlayerCode: req.PlayerCode,
	})
	if playerExist != nil {
		return nil, fmt.Errorf("player code has been used")
	}

	err = r.outport.SavePlayer(ctx, &entity.Player{
		Name:       req.Name,
		PlayerCode: req.PlayerCode,
		PlayerRank: req.PlayerRank,
		CreatedAt:  req.TimeNow,
		UpdatedAt:  req.TimeNow,
		CreatedBy:  "-", // TODO : must change when auth have been created
		UpdatedBy:  "-", // TODO : must change when auth have been created
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
