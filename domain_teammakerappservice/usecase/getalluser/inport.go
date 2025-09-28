package getalluser

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
	Page     int64
	Size     int64
	Username string
	Name     string
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
	Count int64
	Items []*entity.User
}

func (r InportRequest) Validate() error {
	if r.Page <= 0 {

	}
	return nil
}
