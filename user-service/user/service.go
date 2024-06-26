package user

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, id string, password string) (string, error)
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

func (s service) GetUser(ctx context.Context, id string) (*User, error) {
	logger := log.With(s.logger, "method", "GetUser")

	user, err := s.repostory.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("Get user", id)

	return &user, nil
}


func (s service) UpdateUser(ctx context.Context, id string, password string) (string, error) {
	logger := log.With(s.logger, "method", "UpdateUser")

	// Fetch the user first to ensure it exists
	user, err := s.repostory.GetUser(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	// Update the password
	user.Password = password

	// Update the user in the repository
	if err := s.repostory.UpdateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("update user", id)

	return "Success", nil
}
func (s service) DeleteUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteUser")

	// Delete the user from the repository
	if err := s.repostory.DeleteUser(ctx, id); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("delete user", id)

	return "Success", nil
}

