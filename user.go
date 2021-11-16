package gomodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username,required"`
	Email     string             `bson:"email,required"`
	Password  string             `bson:"password,required"`
	CreatedAt time.Time          `bson:"createdAt,required"`
	UpdatedAt time.Time          `bson:"updatedAt,required"`
	DeletedAt time.Time          `bson:"deletedAt,omitempty"`
}

func UserModel(u *User) *User {
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	return u
}

// Update field
