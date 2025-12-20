package runcreatecriteria

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
	Title        string
	TitleAliases []string
	// Rank        []string
	// Rule        []string
	Description string
	Now         time.Time
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
}

func (r InportRequest) Validate() error {
	return nil
}
