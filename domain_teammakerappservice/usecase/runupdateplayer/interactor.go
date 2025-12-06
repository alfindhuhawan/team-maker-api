package runupdateplayer

import (
	"context"
	"fmt"
	"team-maker-api/shared/model/enum"
	"team-maker-api/shared/model/repository"
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
	// playerID, err := primitive.ObjectIDFromHex(req.PlayerID)
	// if err != nil {
	// 	return nil, err
	// }

	playerDataOld, err := r.outport.FindOnePlayer(ctx, enum.IDFilterByEnum, repository.FilterPlayer{
		ID: req.PlayerID,
	})
	if err != nil {
		return nil, err
	}

	// convert user id to string
	// playerID, err := primitive.ObjectIDFromHex(req.PlayerID)
	// if err != nil {
	// 	return nil, fmt.Errorf("invalid player id: %w", err)
	// }

	// must check its from the same id or not
	// fmt.Println("playerDataOld.ID >>")
	// fmt.Println(playerDataOld.ID)
	// fmt.Println("playerID >>")
	// fmt.Println(playerID)
	if playerDataOld.PlayerCode != req.PlayerCode {
		playerExist, _ := r.outport.FindOnePlayer(ctx, enum.PlayerCodeFilterByEnum, repository.FilterPlayer{
			PlayerCode: req.PlayerCode,
		})
		if playerExist != nil {
			return nil, fmt.Errorf("player code has been used")
		}
	}
	// if playerDataOld.ID != playerID {
	// 	fmt.Println("masuk sini")
	// }

	err = r.outport.UpdatePlayer(ctx, req.PlayerID, &repository.UpdatePlayerRequest{
		Name:       req.Name,
		PlayerRank: req.PlayerRank,
		PlayerCode: req.PlayerCode,
		UpdatedAt:  req.TimeNow,
		UpdatedBy:  "-", // TODO : must change when auth have been created
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
