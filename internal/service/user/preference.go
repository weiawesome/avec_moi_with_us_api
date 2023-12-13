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

type ServicePreference struct {
	Repository sql.Repository
}

func (m *ServicePreference) Get(mail string) (responseUser.Preference, error) {
	var response responseUser.Preference
	User := model.User{Mail: mail}
	Preloads := []string{"Genres"}
	if err := m.Repository.PreLoadReadUser(&User, Preloads); errors.As(err, &StatusUtils.ExistSource{}) {
		for _, genre := range User.Genres {
			response.Pairs = append(response.Pairs, responseUser.PreferencePair{Id: genre.GenreId, Value: genre.Val})
		}
		return response, nil
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, responseUtils.ExistError{Message: "Have not sign up"}
	} else {
		return response, err
	}
}
func (m *ServicePreference) GetType() (responseUser.Preference, error) {
	var response responseUser.Preference
	var Genres []model.Genre
	if err := m.Repository.GetGenres(&Genres); err == nil {
		for _, genre := range Genres {
			response.Pairs = append(response.Pairs, responseUser.PreferencePair{Id: genre.GenreId, Value: genre.Val})
		}
		return response, nil
	} else {
		return response, err
	}
}

func (m *ServicePreference) Edit(mail string, preference user.Preference) (responseUser.PreferenceEdit, error) {
	var response responseUser.PreferenceEdit
	if err := m.Repository.UpdateUserGenres(mail, preference.Genres); err == nil {
		return response, nil
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, responseUtils.ExistError{Message: "Have not sign up"}
	} else {
		return response, err
	}
}
