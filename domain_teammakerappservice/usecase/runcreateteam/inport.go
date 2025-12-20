package runcreateteam

import (
	"context"
)

// mirza here

// Inport of Usecase
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase
type InportRequest struct {
	Command string
	// MatchmakingType enum.MatchmakingTypeEnum
	// TotalTeam       int
	// MaxPlayer       int
	// PlayerIDs       []string
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
	Items [][]string
}

func (r InportRequest) Validate() error {
	return nil
}
