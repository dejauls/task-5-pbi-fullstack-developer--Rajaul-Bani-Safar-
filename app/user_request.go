package app

import "github.com/asaskevich/govalidator"

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=6"`
}

func ValidateCreateUserRequest(req CreateUserRequest) error {
	_, err := govalidator.ValidateStruct(req)
	return err
}

func ValidateLoginUserRequest(req LoginUserRequest) error {
	_, err := govalidator.ValidateStruct(req)
	return err
}

func ValidateUpdateUserRequest(req UpdateUserRequest) error {
	_, err := govalidator.ValidateStruct(req)
	return err
}
