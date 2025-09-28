package repository

import (
	"gopkg.in/mgo.v2/bson"
)

type UserSession struct {
	UserId   bson.ObjectId `json:"id"`
	Username string        `json:"username"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Role     string        `json:"role"`
	// DeviceInfo   *entity.DeviceInfo `json:"device_info"`
	MemberCode   string `json:"member_code"`
	AdminHubCode string `json:"admin_hub_code"`
}
