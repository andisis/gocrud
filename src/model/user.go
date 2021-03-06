package model

import "time"

// User represent the users model
type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
