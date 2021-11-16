package gomodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" binding:"required"`
	Username  string             `json:"username,required" binding:"required"`
	Email     string             `json:"email,required" binding:"required"`
	Password  string             `json:"password,required" binding:"required"`
	CreatedAt time.Time          `json:"createdAt,required" binding:"required"`
	UpdatedAt time.Time          `json:"updatedAt,required" binding:"required"`
	DeletedAt time.Time          `json:"deletedAt,omitempty"`
}

func UserModel(u *User) *User {
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	return u
}

// Update field
