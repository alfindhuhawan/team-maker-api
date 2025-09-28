package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Criteria struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Rank        []string      `json:"rank" bson:"rank"`
	Rule        []string      `json:"rule" bson:"rule"`
	Description string        `json:"description" bson:"description"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	CreatedBy   string        `json:"created_by" bson:"created_by"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
	UpdatedBy   string        `json:"updated_by" bson:"updated_by"`
}
