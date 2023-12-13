package user

import (
	"avec_moi_with_us_api/api/request/user"
	responseUser "avec_moi_with_us_api/api/response/user"
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/internal/repository/model"
	"avec_moi_with_us_api/internal/repository/sql"
	StatusUtils "avec_moi_with_us_api/internal/repository/utils"
	"errors"
)

type ServiceEditInformation struct {
	Repository sql.Repository
}

func (m *ServiceEditInformation) Edit(mail string, data user.EditInformation) (responseUser.EditInformation, error) {
	var response responseUser.EditInformation
	User := model.User{Mail: mail}
	if err := m.Repository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		return response, m.Repository.UpdateUser(model.User{Mail: mail, Gender: data.Gender, Name: data.Name})
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, responseUtils.ExistError{Message: "Have not sign up"}
	} else {
		return response, err
	}
}
