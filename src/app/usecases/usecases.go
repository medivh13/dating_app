package usecases

import (
	userUC "github.com/medivh13/dating_app/src/app/usecases/user"
)

type AllUseCases struct {
	UserUC userUC.UserUCInterface
}
