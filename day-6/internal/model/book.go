package model

import "time"

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Writter   string    `json:"writter"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
