package endpoint

import (
	"context"
	service "crud-user/user/service"
	"github.com/go-kit/kit/endpoint"
	schema "crud-user/user/schema"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	UpdateUser endpoint.Endpoint
	DeleteUser endpoint.Endpoint
}

func MakeEndpoints(s service.UserService) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
		UpdateUser: makeUpdateUserEndpoint(s),
		DeleteUser: makeDeleteUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schema.CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return schema.CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schema.GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)

		return schema.GetUserResponse{
			Email: email,
		}, err
	}
}

func makeUpdateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schema.UpdateUserRequest)
		ok, err := s.UpdateUser(ctx, req.Id, req.Email, req.Password)
		return schema.UpdateUserResponse{Ok: ok}, err
	}
}

func makeDeleteUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schema.DeleteUserRequest)
		ok, err := s.DeleteUser(ctx, req.Id)
		return schema.DeleteUserResponse{Ok: ok}, err
	}
}
