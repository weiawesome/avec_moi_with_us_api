package movie

import (
	responseMovie "avec_moi_with_us_api/api/response/movie"
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/internal/repository/redis"
	"avec_moi_with_us_api/internal/repository/sql"
	"avec_moi_with_us_api/internal/repository/utils"
	"errors"
)

type ServiceMovieRecentlyHot struct {
	SqlRepository   sql.Repository
	RedisRepository redis.Repository
}

func (m *ServiceMovieRecentlyHot) Get(page int) (responseMovie.Movies, error) {
	var response responseMovie.Movies
	if movies, err := m.RedisRepository.GetRecentlyHot(page); err == nil {
		response.CurrentPage = page
		response.MovieList = movies
		return response, nil
	} else if errors.As(err, &utils.NotExistSource{}) {
		return response, responseUtils.ExistError{Message: err.Error()}
	} else {
		return response, err
	}
}
