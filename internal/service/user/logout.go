package user

import (
	responseUser "avec_moi_with_us_api/api/response/user"
	"avec_moi_with_us_api/internal/repository/redis"
)

type ServiceLogout struct {
	Repository redis.Repository
}

func (m *ServiceLogout) Logout(jti string) (responseUser.Logout, error) {
	var response responseUser.Logout
	err := m.Repository.SetToBlacklist(jti)

	return response, err
}
