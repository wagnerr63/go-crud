package user

import "go-crud/pkg/exception"

type CreateDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *CreateDTO) Validate() error {
	if dto.Name == "" {
		return exception.NewInvalidRequestError("invalid name")
	}

	if dto.Email == "" {
		return exception.NewInvalidRequestError("invalid email")
	}

	if dto.Password == "" {
		return exception.NewInvalidRequestError("invalid password")
	}

	return nil
}
