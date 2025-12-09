package repository

import (
	"context"
	"team-maker-api/shared/model/entity"
)

type SaveCriteriaRepo interface {
	SaveCriteria(ctx context.Context, obj *entity.Criteria) error
}

type FindAllCriteriaRepo interface {
	FindAllCriteria(ctx context.Context, req FindAllCriteriaRequest) ([]entity.Criteria, int64, error)
}

type FindOneCriteriaRepo interface {
	FindOneCriteria(ctx context.Context, criteriaID string) (*entity.Criteria, error)
}

type DeleteCriteriaRepo interface {
	DeleteCriteria(ctx context.Context, criteriaID string) error
}

type FindAllCriteriaRequest struct {
	Page  int64  `form:"page,omitempty,default=1"`
	Size  int64  `form:"size,omitempty,default=30"`
	Title string `form:"title,omitempty"`
}
