package user

import (
)

type GetUserRequest struct {
	ID    string `form:"id" validate:"omitempty,uuid4"`
	Email string `form:"email" validate:"omitempty,email"`
}

type GetUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
