package user

import (
	"log"

	dto "github.com/medivh13/dating_app/src/app/dto/user"

	userRepo "github.com/medivh13/dating_app/src/infra/persistence/postgres/user"

	helper "github.com/medivh13/dating_app/src/infra/helper"
)

type UserUCInterface interface {
	Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error)
	Login(data *dto.LoginReqDTO) (*dto.LoginRespDTO, error)
}

type userUseCase struct {
	UserRepo userRepo.UserRepository
}

func NewUserUseCase(userRepo userRepo.UserRepository) *userUseCase {
	return &userUseCase{
		UserRepo: userRepo,
	}
}

func (uc *userUseCase) Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error) {
	var resp dto.RegisterRespDTO
	id, err := uc.UserRepo.Register(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp.ID = id
	resp.Token, err = helper.GenerateToken(int(id))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (uc *userUseCase) Login(data *dto.LoginReqDTO) (*dto.LoginRespDTO, error) {
	var resp dto.LoginRespDTO
	result, err := uc.UserRepo.Login(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp.ID = result.ID
	resp.Email = result.Email
	resp.Token, err = helper.GenerateToken(int(result.ID))

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
