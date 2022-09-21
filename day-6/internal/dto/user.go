package dto

import "time"

type UserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name" `
	Email string `json:"email"`
}

type UserPayload struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required" gorm:"unique"`
	Password  string    `json:"password,omitempty" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
