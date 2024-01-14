package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserReqDTOInterface interface {
	ValidateReg() error
	ValidateLogin() error
}
type RegisterReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *RegisterReqDTO) ValidateReg() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Email, validation.Required),
		validation.Field(&dto.Password, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type RegisterRespDTO struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}

type LoginReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *LoginReqDTO) ValidateLogin() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Email, validation.Required),
		validation.Field(&dto.Password, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type LoginRespDTO struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Premium  bool   `json:"premium"`
	Verified bool   `json:"verified"`
	Token    string `json:"token"`
}
