package mock_user

import (
	dto "github.com/medivh13/dating_app/src/app/dto/user"
	models "github.com/medivh13/dating_app/src/infra/models"
	repo "github.com/medivh13/dating_app/src/infra/persistence/postgres/user"

	"github.com/stretchr/testify/mock"
)

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{}
}

var _ repo.UserRepository = &MockUserRepo{}

type MockUserRepo struct {
	mock.Mock
}

func (o *MockUserRepo) Register(data *dto.RegisterReqDTO) (int64, error) {
	args := o.Called(data)

	var (
		err    error
		result int64
	)

	if n, ok := args.Get(0).(int64); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return result, err
}

func (o *MockUserRepo) Login(data *dto.LoginReqDTO) (*models.LoginModel, error) {
	args := o.Called(data)

	var (
		result *models.LoginModel
		err    error
	)

	if n, ok := args.Get(0).(*models.LoginModel); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return result, err
}
