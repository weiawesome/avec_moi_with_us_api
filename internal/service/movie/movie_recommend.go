package movie

import (
	responseMovie "avec_moi_with_us_api/api/response/movie"
	"avec_moi_with_us_api/internal/repository/redis"
	"avec_moi_with_us_api/internal/repository/sql"
)

type ServiceMovieRecommend struct {
	SqlRepository   sql.Repository
	RedisRepository redis.Repository
}

func (m *ServiceMovieRecommend) Get(user string) (responseMovie.Movies, error) {
	var response responseMovie.Movies
	if movies, err := m.SqlRepository.ReadRecommendMovie(user); err == nil {
		response.TotalPages = 1
		response.CurrentPage = 1
		for i := range movies {
			response.MovieList = append(response.MovieList, responseMovie.Movie{MovieId: movies[i].MovieId, Resource: movies[i].Resource, ReleaseYear: movies[i].Year, Title: movies[i].Title, Score: movies[i].Score})
		}
		return response, nil
	} else {
		return response, err
	}
}
