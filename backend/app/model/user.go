package model

import "time"

type User struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"user_name"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
