package getoneuser

import (
	"context"
	"team-maker-api/shared/model/enum"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockery --name Outport -output mocks/

type getOneUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getOneUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getOneUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	// Check username first
	userID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return nil, err
	}

	dataUser, err := r.outport.FindOneUser(ctx, enum.IDFilterByEnum, userID)
	if err != nil {
		return nil, err
	}

	res.Item = dataUser

	return res, nil
}
