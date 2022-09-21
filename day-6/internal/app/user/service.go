package user

import (
	"alterra-agmc/day-6/internal/dto"
	"alterra-agmc/day-6/internal/factory"
	"alterra-agmc/day-6/internal/repository"
	res "alterra-agmc/day-6/pkg/utill/response"
	"context"
	"time"

	"alterra-agmc/day-6/internal/model"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	CreateUser(ctx context.Context, payload *dto.UserPayload) error
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) CreateUser(ctx context.Context, payload *dto.UserPayload) error {
	var userCreate = model.User{
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.UserRepository.CreateUser(ctx, userCreate)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return nil
}
