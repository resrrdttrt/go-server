package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, email string, name string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
	UpdateUser(ctx context.Context, id string, name string, password string) (string, error)
	DeleteUser(ctx context.Context, id string) (string, error)
}



type service struct {
	repostory UserRepository
	logger    log.Logger
}

func NewService(rep UserRepository, logger log.Logger) UserService {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")
	
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if err := s.repostory.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)

	return "Success", nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repostory.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get user", id)

	return email, nil
}
