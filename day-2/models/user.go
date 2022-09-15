package models

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
