package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	ID           bson.ObjectId `json:"id" bson:"_id"`
	PlayerCode   string        `json:"player_code" bson:"player_code"`
	Name         string        `json:"name" bson:"name"`
	CriteriaRank []string      `json:"criteria_rank" bson:"criteria_rank"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
	CreatedBy    string        `json:"created_by" bson:"created_by"`
	UpdatedAt    time.Time     `json:"updated_at" bson:"updated_at"`
	UpdatedBy    string        `json:"updated_by" bson:"updated_by"`
}
