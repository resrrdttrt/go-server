package schemas

import (
)

type GetUserRequest struct {
	ID    string `form:"id" validate:"omitempty,uuid4"`
}

type GetUserResponse struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
