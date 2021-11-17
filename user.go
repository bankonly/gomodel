package gomodel

import (
	"time"
)

type User struct {
	Username  string    `json:"username" bson:"username" valid:"required~ERR4001"`
	Email     string    `json:"email" bson:"email" valid:"required~ERR4002"`
	Password  string    `json:"password" bson:"password" valid:"required~ERR4003"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" bson:"deletedAt,omitempty"`
}

func UserModel(u User) User {
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	return u
}

// Update field
