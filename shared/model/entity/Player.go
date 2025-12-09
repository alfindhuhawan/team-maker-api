package entity

import (
	"team-maker-api/shared/model/enum"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID         primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	PlayerCode string              `json:"player_code" bson:"player_code"`
	Name       string              `json:"name" bson:"name"`
	PlayerRank enum.PlayerRankEnum `json:"player_rank" bson:"player_rank"`
	Criteria   []string            `json:"criteria" bson:"criteria"` // TODO : next step will be used more detail criteria
	CreatedAt  time.Time           `json:"created_at" bson:"created_at"`
	CreatedBy  string              `json:"created_by" bson:"created_by"`
	UpdatedAt  time.Time           `json:"updated_at" bson:"updated_at"`
	UpdatedBy  string              `json:"updated_by" bson:"updated_by"`
}
