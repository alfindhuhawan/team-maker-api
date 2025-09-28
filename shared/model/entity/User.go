package entity

import (
	"team-maker-api/shared/model/enum"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// struct address dari CMSrevamp API
type User struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string        `json:"username" bson:"username,omitempty"`
	Password  string        `json:"password" bson:"password,omitempty"`
	Name      string        `json:"name" bson:"name,omitempty"`
	Role      enum.RoleEnum `json:"role" bson:"role,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at,omitempty"`
	CreatedBy string        `json:"created_by" bson:"created_by,omitempty"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at,omitempty"`
	UpdatedBy string        `json:"updated_by" bson:"updated_by,omitempty"`
}
