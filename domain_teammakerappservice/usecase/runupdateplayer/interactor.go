package runupdateplayer

import (
	"context"
	"team-maker-api/shared/model/enum"
	"team-maker-api/shared/model/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockery --name Outport -output mocks/

type runUpdatePlayerInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runUpdatePlayerInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runUpdatePlayerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	// Check username first
	playerID, err := primitive.ObjectIDFromHex(req.PlayerID)
	if err != nil {
		return nil, err
	}

	_, err = r.outport.FindOnePlayer(ctx, enum.IDFilterByEnum, playerID)
	if err != nil {
		return nil, err
	}

	err = r.outport.UpdatePlayer(ctx, playerID, &repository.UpdatePlayerRequest{
		Name:       req.Name,
		PlayerRank: string(req.PlayerRank),
		UpdatedAt:  req.TimeNow,
		UpdatedBy:  "-", // TODO : must change when auth have been created
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
