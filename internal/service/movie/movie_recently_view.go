package movie

import (
	responseMovie "avec_moi_with_us_api/api/response/movie"
	"avec_moi_with_us_api/internal/repository/redis"
	"avec_moi_with_us_api/internal/repository/sql"
)

type ServiceMovieRecentlyView struct {
	SqlRepository   sql.Repository
	RedisRepository redis.Repository
}

func (m *ServiceMovieRecentlyView) Get(user string, page int, pageSize int) (responseMovie.Movies, error) {
	var response responseMovie.Movies
	if movies, totalPages, err := m.SqlRepository.ReadRecentlyViewMovieByPage(user, page, pageSize); err == nil {
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
