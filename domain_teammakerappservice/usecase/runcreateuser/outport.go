package runcreateuser

import "team-maker-api/shared/model/repository"

// Outport of usecase
type Outport interface {
	repository.SaveUserRepo
	repository.FindOneUserRepo
}
