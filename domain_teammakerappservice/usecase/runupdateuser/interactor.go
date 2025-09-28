package runupdateuser

import (
	"context"
	"fmt"
	"team-maker-api/shared/infrastructure/util"
	"team-maker-api/shared/model/enum"
	"team-maker-api/shared/model/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockery --name Outport -output mocks/

type runUpdateUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runUpdateUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runUpdateUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	oldDataUser, err := r.outport.FindOneUser(ctx, enum.IDFilterByEnum, userID)
	if err != nil {
		return nil, err
	}

	userFound, _ := r.outport.FindOneUser(ctx, enum.UsernameFilterByEnum, req.Username)
	if userFound != nil {
		if oldDataUser.Username != req.Username {
			return nil, fmt.Errorf("username telah digunakan")
		}
	}

	// Insert user
	password, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	fmt.Println("req.Name >>")
	fmt.Println(req.Name)

	err = r.outport.UpdateUser(ctx, userID, &repository.UpdateRequest{
		Username:  req.Username,
		Password:  password,
		Name:      req.Name,
		Role:      oldDataUser.Role,
		UpdatedAt: req.TimeNow,
		UpdatedBy: "-", // TODO : must change when auth have been created
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
