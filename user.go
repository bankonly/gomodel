package gomodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
	DeletedAt time.Time          `bson:"deletedAt,omitempty"`
}

func UserModel(username, email, password string) *User {
	u := User{}
	u.Username = username
	u.Email = email
	u.Password = password
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	return &u
}

// Update field
