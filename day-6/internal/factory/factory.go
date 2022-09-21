package factory

import (
	"alterra-agmc/day-6/database"
	"alterra-agmc/day-6/internal/repository"
)

type Factory struct {
	UserRepository repository.User
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		UserRepository: repository.NewUser(db),
	}
}
