package schema
import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	GetUserRequest struct {
		Id string `json:"id"`
	}
	GetUserResponse struct {
		Email string `json:"email"`
	}

	UpdateUserRequest struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UpdateUserResponse struct {
		Ok string `json:"ok"`
	}

	DeleteUserRequest struct {
		Id string `json:"id"`
	}
	DeleteUserResponse struct {
		Ok string `json:"ok"`
	}
)



func DecodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error){
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return CreateUserRequest{}, err
	}
	return req, nil
}


func DecodeGetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)
	req = GetUserRequest{
		Id: vars["id"],
	}
	return req, nil
}


func DecodeUpdateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return UpdateUserRequest{}, err
	}
	return req, nil
}

func DecodeDeleteUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req DeleteUserRequest
	vars := mux.Vars(r)
	req = DeleteUserRequest{
		Id: vars["id"],
	}
	return req, nil
}


func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

