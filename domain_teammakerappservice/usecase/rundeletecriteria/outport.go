package rundeletecriteria

import "team-maker-api/shared/model/repository"

// Outport of usecase
type Outport interface {
	repository.FindOneCriteriaRepo
	repository.DeleteCriteriaRepo
}
