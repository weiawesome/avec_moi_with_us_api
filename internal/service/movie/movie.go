package movie

import (
	responseMovie "avec_moi_with_us_api/api/response/movie"
	"avec_moi_with_us_api/api/utils"
	"avec_moi_with_us_api/internal/repository/redis"
	"avec_moi_with_us_api/internal/repository/sql"
	"avec_moi_with_us_api/internal/service/logger"
)

type ServiceMovie struct {
	SqlRepository   sql.Repository
	RedisRepository redis.Repository
}

func (m *ServiceMovie) Get(page int, pageSize int, randomSeed int) (responseMovie.Movies, error) {
	var response responseMovie.Movies
	cache, err := m.RedisRepository.LoadMovieCache(page, randomSeed)
	if err == nil {
		return cache, nil
	}
	if movies, totalPages, err := m.SqlRepository.ReadMovieByPage(page, pageSize, randomSeed); err == nil {
		response.TotalPages = totalPages
		response.CurrentPage = page
		for i := range movies {
			response.MovieList = append(response.MovieList, responseMovie.Movie{MovieId: movies[i].MovieId, Resource: movies[i].Resource, ReleaseYear: movies[i].Year, Title: movies[i].Title, Score: movies[i].Score})
		}
		go func() {
			err := m.RedisRepository.SaveMovieCache(page, randomSeed, response)
			if err != nil {
				serviceLogger := logger.ServiceLogger{Repository: m.RedisRepository}
				serviceLogger.Error(utils.LogData{Event: "Failure to save cache", User: "system", Details: err.Error()})
			}
		}()
		return response, nil
	} else {
		return response, err
	}
}
