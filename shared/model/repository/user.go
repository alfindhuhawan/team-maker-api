package repository

import (
	"context"
	"team-maker-api/shared/model/entity"
	"team-maker-api/shared/model/enum"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
}

type UpdateUserRepo interface {
	UpdateUser(ctx context.Context, userID primitive.ObjectID, obj *UpdateRequest) error
}

type FindOneUserRepo interface {
	FindOneUser(ctx context.Context, filterBy enum.FilterByEnum, filter interface{}) (*entity.User, error)
}

type FindAllUserRepo interface {
	FindAllUser(ctx context.Context, req FindAllUserRequest) ([]*entity.User, int64, error)
}

type FindAllUserRequest struct {
	Page     int64  `form:"page,omitempty,default=1"`
	Size     int64  `form:"size,omitempty,default=30"`
	Name     string `form:"name,omitempty"`
	Username string `form:"username,omitempty"`
}

type UpdateRequest struct {
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"password" bson:"password"`
	Name      string        `json:"name" bson:"name"`
	Role      enum.RoleEnum `json:"role" bson:"role"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	UpdatedBy string        `json:"updated_by" bson:"updated_by"`
}
