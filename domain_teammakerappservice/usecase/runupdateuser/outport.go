package runupdateuser

import "team-maker-api/shared/model/repository"

// Outport of usecase
type Outport interface {
	repository.UpdateUserRepo
	repository.FindOneUserRepo
}
