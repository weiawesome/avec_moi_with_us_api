package movie

import (
	responseMovie "avec_moi_with_us_api/api/response/movie"
	"avec_moi_with_us_api/api/utils"
	"avec_moi_with_us_api/internal/repository/chroma"
	"avec_moi_with_us_api/internal/repository/redis"
	"avec_moi_with_us_api/internal/repository/sql"
	"avec_moi_with_us_api/internal/service/logger"
)

type ServiceMovieSearch struct {
	SqlRepository    sql.Repository
	RedisRepository  redis.Repository
	ChromaRepository chroma.Repository
}

func (m *ServiceMovieSearch) Search(content string, page int, pageSize int) (responseMovie.Movies, error) {
	cache, err := m.RedisRepository.LoadSearchCache(content, page)
	if err == nil {
		return cache, err
	}
	search, err := m.RedisRepository.LoadSearchCandidateCache(content)
	if err != nil {
		search, _ = m.ChromaRepository.Search(content)
		go func() {
			err := m.RedisRepository.SaveSearchCandidateCache(content, search)
			if err != nil {
				serviceLogger := logger.ServiceLogger{Repository: m.RedisRepository}
				serviceLogger.Error(utils.LogData{Event: "Failure to save cache", User: "system", Details: err.Error()})
			}
		}()
	}
	var response responseMovie.Movies
	if movies, totalPages, err := m.SqlRepository.ReadSearchMovieByPage(search, content, page, pageSize); err == nil {
		response.TotalPages = totalPages
		response.CurrentPage = page
		for i := range movies {
			response.MovieList = append(response.MovieList, responseMovie.Movie{MovieId: movies[i].MovieId, Resource: movies[i].Resource, ReleaseYear: movies[i].Year, Title: movies[i].Title, Score: movies[i].Score})
		}
		go func() {
			err := m.RedisRepository.SaveSearchCache(content, page, response)
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
