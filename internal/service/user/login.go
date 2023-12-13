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

type ServiceLogin struct {
	Repository sql.Repository
}

func (m *ServiceLogin) Login(data user.Login) (responseUser.Login, error) {
	var response responseUser.Login
	User := model.User{Mail: data.Mail}

	if err := m.Repository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		if utils.ValidatePassword(data.Password, User.Password, User.Salt) {
			jwt, _, err := utils.GetJwt(data.Mail)
			return responseUser.Login{Token: jwt}, err
		} else {
			return response, responseUtils.ExistError{Message: "Password error"}
		}
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, responseUtils.ExistError{Message: "Have not sign up"}
	} else {
		return response, err
	}
}
