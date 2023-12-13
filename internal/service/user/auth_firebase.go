package user

import (
	responseUser "avec_moi_with_us_api/api/response/user"
	"avec_moi_with_us_api/api/utils"
	"avec_moi_with_us_api/internal/repository/firebase"
	"avec_moi_with_us_api/internal/repository/model"
	"avec_moi_with_us_api/internal/repository/sql"
	StatusUtils "avec_moi_with_us_api/internal/repository/utils"
	"errors"
)

type ServiceAuthFirebase struct {
	SqlRepository      sql.Repository
	FirebaseRepository firebase.Repository
}

func (m *ServiceAuthFirebase) Auth(IdToken string) (responseUser.Login, error) {
	var response responseUser.Login
	validate, err := m.FirebaseRepository.Validate(IdToken)
	if err != nil {
		return responseUser.Login{}, err
	}
	User := model.User{Mail: validate.Mail}
	if err := m.SqlRepository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		jwt, _, err := utils.GetJwt(validate.Mail)
		return responseUser.Login{Token: jwt}, err
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		err = m.SqlRepository.CreateUser(model.User{Name: validate.Name, Mail: validate.Mail, Gender: validate.Gender})
		if err != nil {
			return response, err
		}
		jwt, _, err := utils.GetJwt(validate.Mail)
		return responseUser.Login{Token: jwt}, err
	} else {
		return response, err
	}
}
