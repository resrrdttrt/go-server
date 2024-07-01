package endpoint

import (
	"context"
	service "crud-user/user/service"
	"github.com/go-kit/kit/endpoint"
	schema "crud-user/user/schema"
)

type GWEndpoints struct {
	GetUser    endpoint.Endpoint
	UpdateUser endpoint.Endpoint
}

func MakeGWEndpoints(s service.GatewayService) GWEndpoints {
	return GWEndpoints{
		GetUser:    makeGetGWEndpoint(s),
		UpdateUser: makeUpdateGWEndpoint(s),
	}
}


func makeGetGWEndpoint(s service.GatewayService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schema.GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)

		return schema.GetUserResponse{
			Email: email,
		}, err
	}
}

func makeUpdateGWEndpoint(s service.GatewayService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schema.UpdateUserRequest)
		ok, err := s.UpdateUser(ctx, req.Id, req.Email, req.Password)
		return schema.UpdateUserResponse{Ok: ok}, err
	}
}

