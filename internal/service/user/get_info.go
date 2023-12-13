package user

import (
	responseUser "avec_moi_with_us_api/api/response/user"
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/internal/repository/model"
	"avec_moi_with_us_api/internal/repository/sql"
	StatusUtils "avec_moi_with_us_api/internal/repository/utils"
	"errors"
)

type ServiceGetInfo struct {
	Repository sql.Repository
}

func (m *ServiceGetInfo) Get(mail string) (responseUser.Information, error) {
	var response responseUser.Information
	User := model.User{Mail: mail}

	if err := m.Repository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		return responseUser.Information{Gender: User.Gender, Name: User.Name, Mail: User.Mail}, nil
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, responseUtils.ExistError{Message: "Have not sign up"}
	} else {
		return response, err
	}
}
