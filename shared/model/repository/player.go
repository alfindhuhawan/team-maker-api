package repository

import (
	"context"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/enum"
	"time"
)

type SavePlayerRepo interface {
	SavePlayer(ctx context.Context, obj *entity.Player) error
}

type UpdatePlayerRepo interface {
	UpdatePlayer(ctx context.Context, userID string, obj *UpdatePlayerRequest) error
}

type FindOnePlayerRepo interface {
	FindOnePlayer(ctx context.Context, filterBy enum.FilterByEnum, filterPlayer FilterPlayer) (*entity.Player, error)
}

type FindAllPlayerRepo interface {
	FindAllPlayer(ctx context.Context, req FindAllPlayerRequest) ([]*entity.Player, int64, error)
}

type FindAllPlayerListRepo interface {
	FindAllPlayerList(ctx context.Context, req *FindAllPlayerListRequest) ([]*entity.Player, error)
}

type DeletePlayerRepo interface {
	DeletePlayer(ctx context.Context, userID string) error
}

type FindAllPlayerRequest struct {
	Page       int64               `form:"page,omitempty,default=1"`
	Size       int64               `form:"size,omitempty,default=30"`
	Name       string              `form:"name,omitempty"`
	PlayerRank enum.PlayerRankEnum `form:"player_rank,omitempty"`
}

type UpdatePlayerRequest struct {
	Name       string              `json:"name" bson:"name"`
	PlayerRank enum.PlayerRankEnum `json:"player_rank" bson:"player_rank"`
	PlayerCode string              `json:"player_code" bson:"player_code"`
	UpdatedAt  time.Time           `json:"updated_at" bson:"updated_at"`
	UpdatedBy  string              `json:"updated_by" bson:"updated_by"`
}

type FindAllPlayerListRequest struct {
	PlayerIDs []string
}

type FilterPlayer struct {
	ID         string
	Name       string
	PlayerCode string
}
