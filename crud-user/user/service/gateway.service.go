package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	pb "crud-user/user/pb"
)

type GatewayService interface {
	GetUser(ctx context.Context, id string) (string, error)
	UpdateUser(ctx context.Context, id string, email string, password string) (string, error)
}

type gatewayservice struct {
	grpc_client pb.UserServiceClient
	logger    log.Logger
}

func NewGatewayService(grpc_client pb.UserServiceClient, logger log.Logger) GatewayService {
	return &gatewayservice{
		grpc_client: grpc_client,
		logger:    logger,
	}
}


func (s gatewayservice) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")
	req := &pb.GetUserRequest{UserId: id}
	res, err := s.grpc_client.GetUser(ctx, req)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get user", id)

	return res.User.Email, nil
}

func (s gatewayservice) UpdateUser(ctx context.Context, id string, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "UpdateUser")

	user := &pb.User{
		Id:       "194fc53b-c3ff-4aaa-867e-348ddd490d47",
		Email:    "newemail@example.com",
		Password: "newpassword",
	}
	r, err := s.grpc_client.UpdateUser(ctx, &pb.UpdateUserRequest{
		User: user,
	})
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("update user", id)

	if r.Success {
		return "Success", nil
	} else {
		return "Fail", nil
	}
}

