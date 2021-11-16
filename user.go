package gomodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" binding:"required"`
	Username  string             `bson:"username,required" binding:"required"`
	Email     string             `bson:"email,required" binding:"required"`
	Password  string             `bson:"password,required" binding:"required"`
	CreatedAt time.Time          `bson:"createdAt,required" binding:"required"`
	UpdatedAt time.Time          `bson:"updatedAt,required" binding:"required"`
	DeletedAt time.Time          `bson:"deletedAt,omitempty"`
}

func UserModel(u *User) *User {
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	return u
}

// Update field
