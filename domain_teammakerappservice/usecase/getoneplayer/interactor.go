package getoneplayer

import (
	"context"
	"team-maker-api/shared/model/enum"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockery --name Outport -output mocks/

type getOnePlayerInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getOnePlayerInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getOnePlayerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	dataUser, err := r.outport.FindOnePlayer(ctx, enum.IDFilterByEnum, playerID)
	if err != nil {
		return nil, err
	}

	res.Item = dataUser

	return res, nil
}
