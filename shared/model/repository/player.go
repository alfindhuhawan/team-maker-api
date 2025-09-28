package repository

import (
	"context"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/enum"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SavePlayerRepo interface {
	SavePlayer(ctx context.Context, obj *entity.Player) error
}

type UpdatePlayerRepo interface {
	UpdatePlayer(ctx context.Context, userID primitive.ObjectID, obj *UpdatePlayerRequest) error
}

type FindOnePlayerRepo interface {
	FindOnePlayer(ctx context.Context, filterBy enum.FilterByEnum, filter interface{}) (*entity.Player, error)
}

type FindAllPlayerRepo interface {
	FindAllPlayer(ctx context.Context, req FindAllPlayerRequest) ([]*entity.Player, int64, error)
}

type FindAllPlayerListRepo interface {
	FindAllPlayerList(ctx context.Context, req *FindAllPlayerListRequest) ([]*entity.Player, error)
}

type FindAllPlayerRequest struct {
	Page int64  `form:"page,omitempty,default=1"`
	Size int64  `form:"size,omitempty,default=30"`
	Name string `form:"name,omitempty"`
}

type UpdatePlayerRequest struct {
	Name       string    `json:"name" bson:"name"`
	PlayerRank string    `json:"player_rank" bson:"player_rank"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy  string    `json:"updated_by" bson:"updated_by"`
}

type FindAllPlayerListRequest struct {
	PlayerIDs []string
}
