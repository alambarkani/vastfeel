package models

import "time"

type Role string

const (
	Admin Role = "admin"
	Guest Role = "guest"
)

func (role Role) IsValidRole() bool {
	switch role {
	case Admin, Guest:
		return true
	default:
		return false
	}
}

type User struct {
	ID        int    	`json:"id" form:"id"`
	Username  string 	`json:"username" form:"username" validate:"required"`
	Email     string 	`json:"email" form:"email" validate:"required,email"`
	Password  string 	`json:"password" form:"password" validate:"required,min=8"`
	Role      Role   	`json:"role" form:"role" validate:"required"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}