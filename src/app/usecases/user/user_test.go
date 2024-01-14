package user

import (
	"errors"

	mockDTO "github.com/medivh13/dating_app/mocks/app/dto/user"
	mockRepo "github.com/medivh13/dating_app/mocks/infra/persistence/postgres/user"

	"testing"

	dto "github.com/medivh13/dating_app/src/app/dto/user"
	model "github.com/medivh13/dating_app/src/infra/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseList struct {
	suite.Suite
	mockDTO  *mockDTO.MockUserDTO
	mockRepo *mockRepo.MockUserRepo
	useCase  UserUCInterface

	dtoRegister  *dto.RegisterReqDTO
	dtoLogin     *dto.LoginReqDTO
	loginModel *model.LoginModel
}

func (suite *UserUseCaseList) SetupTest() {

	suite.mockDTO = new(mockDTO.MockUserDTO)
	suite.mockRepo = new(mockRepo.MockUserRepo)
	suite.useCase = NewUserUseCase(suite.mockRepo)

	suite.dtoRegister = &dto.RegisterReqDTO{
		Email:    "email@gmail.com",
		Password: "1234567",
	}

	suite.dtoLogin = &dto.LoginReqDTO{
		Email:    "email@gmail.com",
		Password: "1234567",
	}

	suite.loginModel = &model.LoginModel{
		ID:       1,
		Email:    "email",
		Premium:  false,
		Verified: false,
		Password: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MDUyMDk0MTF9.bxcz045FQVaW2PICO0GAKCT7KYLdxxVmByc2bE9_Zso",
		
	}
}

func (u *UserUseCaseList) TestRegisterSuccess() {
	u.mockRepo.Mock.On("Register", mock.Anything).Return(mock.Anything, nil)
	_, err := u.useCase.Register(u.dtoRegister)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestRegisterFail() {
	u.mockRepo.Mock.On("Register", mock.Anything).Return(0, errors.New(mock.Anything))
	_, err := u.useCase.Register(u.dtoRegister)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *UserUseCaseList) TestLoginSuccess() {
	u.mockRepo.Mock.On("Login", mock.Anything).Return(u.loginModel, nil)
	_, err := u.useCase.Login(u.dtoLogin)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestLoginFail() {
	u.mockRepo.Mock.On("Login", mock.Anything).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.Login(u.dtoLogin)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UserUseCaseList))
}
