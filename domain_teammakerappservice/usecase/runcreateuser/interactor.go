package runcreateuser

import (
	"context"
	"fmt"
	"team-maker-api/shared/infrastructure/util"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/enum"
)

//go:generate mockery --name Outport -output mocks/

type runCreateUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runCreateUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runCreateUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	// Check username first
	userFound, _ := r.outport.FindOneUser(ctx, enum.UsernameFilterByEnum, req.Username)
	if userFound != nil {
		return nil, fmt.Errorf("username telah digunakan")
	}

	// Insert user
	password, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveUser(ctx, &entity.User{
		Username:  req.Username,
		Password:  password,
		Name:      req.Name,
		Role:      enum.AdminRoleEnum,
		CreatedAt: req.TimeNow,
		UpdatedAt: req.TimeNow,
		CreatedBy: "-", // TODO : must change when auth have been created
		UpdatedBy: "-", // TODO : must change when auth have been created
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
