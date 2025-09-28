package getoneplayer

import (
	"context"
	"team-maker-api/shared/model/entity"
)

// mirza here

// Inport of Usecase
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase
type InportRequest struct {
	PlayerID string
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
	Item *entity.Player
}

func (r InportRequest) Validate() error {
	return nil
}
