package repository

import (
	"alterra-agmc/day-6/internal/model"
	"context"

	"gorm.io/gorm"
)

type User interface {
	// GetUser() (*[]model.User, error)
	// GetUserByID(int) (*model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	// UpdateUser(int, model.User) error
	// DeleteUser(int) (*model.User, error)
}

type user struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{db}
}

func (p *user) CreateUser(ctx context.Context, data model.User) error {
	return p.Db.WithContext(ctx).Model(&model.User{}).Create(&data).Error
}
