package user

import (
	"avec_moi_with_us_api/api/request/user"
	responseUser "avec_moi_with_us_api/api/response/user"
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"avec_moi_with_us_api/internal/repository/model"
	"avec_moi_with_us_api/internal/repository/sql"
	StatusUtils "avec_moi_with_us_api/internal/repository/utils"
	"errors"
)

type ServiceEditPassword struct {
	Repository sql.Repository
}

func (m *ServiceEditPassword) Edit(mail string, data user.EditPassword) (responseUser.EditPassword, error) {
	var response responseUser.EditPassword
	User := model.User{Mail: mail}

	if err := m.Repository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		if utils.ValidatePassword(data.CurrentPassword, User.Password, User.Salt) {
			salt, err := utils.GenerateSalt()
			if err != nil {
				return response, err
			}
			password := utils.GenerateHashPassword(data.EditPassword, salt)
			return response, m.Repository.UpdateUser(model.User{Mail: mail, Password: password, Salt: salt})
		} else {
			return response, responseUtils.ExistError{Message: "Password error"}
		}
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, responseUtils.ExistError{Message: "Have not sign up"}
	} else {
		return response, err
	}
}
