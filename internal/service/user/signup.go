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

type ServiceSignup struct {
	Repository sql.Repository
}

func (m *ServiceSignup) Signup(data user.SignUp) (responseUser.Signup, error) {
	var response responseUser.Signup

	mail := data.Mail
	password := data.Password
	gender := data.Gender
	name := data.Name

	User := model.User{Mail: mail}
	if err := m.Repository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		return response, responseUtils.ExistError{Message: "Have been registered"}
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		salt, err := utils.GenerateSalt()
		if err != nil {
			return response, err
		}
		password = utils.GenerateHashPassword(password, salt)
		err = m.Repository.CreateUser(model.User{Name: name, Mail: mail, Gender: gender, Password: password, Salt: salt})
		return response, err
	} else {
		return response, err
	}
}
