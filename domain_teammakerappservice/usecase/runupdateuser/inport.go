package runupdateuser

import (
	"context"
	"time"
)

// mirza here

// Inport of Usecase
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase
type InportRequest struct {
	UserID   string
	Username string
	Password string
	Name     string
	TimeNow  time.Time
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
}

func (r InportRequest) Validate() error {
	return nil
}
