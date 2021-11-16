package gomodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Username  string             `bson:"username" json:"username" validate:"required"`
	Email     string             `bson:"email" json:"email" validate:"required"`
	Password  string             `bson:"password" json:"password" validate:"required"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	DeletedAt time.Time          `bson:"deletedAt,omitempty" json:"deletedAt"`
}

func UserModel(u *User) *User {
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	return u
}

// Update field
