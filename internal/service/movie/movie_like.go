package movie

import (
	responseMovie "avec_moi_with_us_api/api/response/movie"
	"avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/internal/repository/model"
	"avec_moi_with_us_api/internal/repository/redis"
	"avec_moi_with_us_api/internal/repository/sql"
	StatusUtils "avec_moi_with_us_api/internal/repository/utils"
	"errors"
)

type ServiceMovieLike struct {
	SqlRepository   sql.Repository
	RedisRepository redis.Repository
}

func (m *ServiceMovieLike) Get(user string, page int, pageSize int) (responseMovie.Movies, error) {
	var response responseMovie.Movies
	if movies, totalPages, err := m.SqlRepository.ReadLikeMovieByPage(user, page, pageSize); err == nil {
		response.TotalPages = totalPages
		response.CurrentPage = page
		for i := range movies {
			response.MovieList = append(response.MovieList, responseMovie.Movie{MovieId: movies[i].MovieId, Resource: movies[i].Resource, ReleaseYear: movies[i].Year, Title: movies[i].Title, Score: movies[i].Score})
		}
		return response, nil
	} else {
		return response, err
	}
}

func (m *ServiceMovieLike) IsLike(user string, id string) (responseMovie.MovieIsLike, error) {
	var response responseMovie.MovieIsLike
	userLikeMovie := model.UserLikeMovie{UserMail: user, MovieId: id}
	if err := m.SqlRepository.ReadLikeMovie(&userLikeMovie); errors.As(err, &StatusUtils.ExistSource{}) {
		return response, nil
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, utils.ExistError{Message: "Source not exist"}
	} else {
		return response, err
	}
}
func (m *ServiceMovieLike) Like(user string, id string) (responseMovie.LikeMovieResult, error) {
	var response responseMovie.LikeMovieResult
	User := model.User{Mail: user}
	Movie := model.Movie{MovieId: id}
	if err := m.SqlRepository.ReadMovie(&Movie); errors.As(err, &StatusUtils.ExistSource{}) {
		if userErr := m.SqlRepository.ReadUser(&User); errors.As(userErr, &StatusUtils.ExistSource{}) {
			err := m.SqlRepository.CreateOrUpdateUserLikeMovie(model.UserLikeMovie{UserMail: user, MovieId: id})
			return responseMovie.LikeMovieResult{}, err
		} else if errors.As(userErr, &StatusUtils.NotExistSource{}) {
			return response, utils.ExistError{Message: "User not exist"}
		} else {
			return response, userErr
		}
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, utils.ExistError{Message: "Source not exist"}
	} else {
		return response, err
	}
}
func (m *ServiceMovieLike) DisLike(user string, id string) (responseMovie.LikeMovieResult, error) {
	var response responseMovie.LikeMovieResult
	User := model.User{Mail: user}
	Movie := model.Movie{MovieId: id}
	if err := m.SqlRepository.ReadMovie(&Movie); errors.As(err, &StatusUtils.ExistSource{}) {
		if userErr := m.SqlRepository.ReadUser(&User); errors.As(userErr, &StatusUtils.ExistSource{}) {
			if err := m.SqlRepository.DeleteMovieLike(model.UserLikeMovie{UserMail: user, MovieId: id}); err == nil {
				return response, nil
			} else if errors.As(err, &StatusUtils.NotExistSource{}) {
				return response, utils.ExistError{Message: "Like resource not exist"}
			} else {
				return response, userErr
			}
		} else if errors.As(userErr, &StatusUtils.NotExistSource{}) {
			return response, utils.ExistError{Message: "User not exist"}
		} else {
			return response, userErr
		}
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, utils.ExistError{Message: "Source not exist"}
	} else {
		return response, err
	}
}
