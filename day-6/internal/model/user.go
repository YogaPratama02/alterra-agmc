package model

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required" gorm:"unique"`
	Password  string    `json:"password,omitempty" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Token struct {
	Token string `json:"token"`
}
