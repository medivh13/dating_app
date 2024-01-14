package mock_user

import (
	dto "github.com/medivh13/dating_app/src/app/dto/user"

	"github.com/stretchr/testify/mock"
)

type MockUserDTO struct {
	mock.Mock
}

func NewMockUserDTO() *MockUserDTO {
	return &MockUserDTO{}
}

var _ dto.UserReqDTOInterface = &MockUserDTO{}

func (m *MockUserDTO) ValidateReg() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}

func (m *MockUserDTO) ValidateLogin() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
