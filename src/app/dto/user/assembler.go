package user

import (
	model "github.com/medivh13/dating_app/src/infra/models"
)

func ToReturnUserLogin(d *model.LoginModel) *LoginRespDTO {
	return &LoginRespDTO{
		ID:       d.ID,
		Email:    d.Email,
		Premium:  d.Premium,
		Verified: d.Premium,
	}
}
