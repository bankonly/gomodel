package gomodel

import "time"

type User struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt bool      `json:"deletedAt"`
}

func UserModel(username, email, password string) *User {
	u := User{}
	u.Username = username
	u.Email = email
	u.Password = password
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	u.DeletedAt = false
	return &u
}

// Update field
