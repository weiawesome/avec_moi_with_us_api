package movie

import (
	responseMovie "avec_moi_with_us_api/api/response/movie"
	"avec_moi_with_us_api/api/response/utils"
	logUtils "avec_moi_with_us_api/api/utils"
	"avec_moi_with_us_api/internal/repository/model"
	"avec_moi_with_us_api/internal/repository/redis"
	"avec_moi_with_us_api/internal/repository/sql"
	StatusUtils "avec_moi_with_us_api/internal/repository/utils"
	"avec_moi_with_us_api/internal/service/logger"
	"errors"
	"time"
)

type ServiceMovieSpecific struct {
	SqlRepository   sql.Repository
	RedisRepository redis.Repository
}

func (m *ServiceMovieSpecific) Get(id string) (responseMovie.SpecificMovie, error) {
	cache, err := m.RedisRepository.LoadSpecificCache(id)
	if err == nil {
		return cache, err
	}
	var response responseMovie.SpecificMovie
	Movie := model.Movie{MovieId: id}
	Preloads := []string{"Actors", "Directors", "Genres", "MovieTitle", "MovieRank"}
	if err := m.SqlRepository.PreLoadReadMovie(&Movie, Preloads); errors.As(err, &StatusUtils.ExistSource{}) {
		response = responseMovie.SpecificMovie{
			MovieId:       Movie.MovieId,
			Resource:      Movie.Resource,
			ReleaseYear:   Movie.Year,
			Title:         Movie.Title,
			OriginalTitle: Movie.OriginalTitle,
			Introduction:  Movie.Introduction,
		}
		for _, title := range Movie.MovieTitle {
			response.Titles = append(response.Titles, title.Title)
		}
		for _, genre := range Movie.Genres {
			response.Categories = append(response.Categories, genre.Val)
		}
		for _, actor := range Movie.Actors {
			response.Actors = append(response.Actors, responseMovie.Celebrity{Name: actor.Name, Resource: actor.Resource, Gender: actor.Gender})
		}
		for _, director := range Movie.Directors {
			response.Directors = append(response.Directors, responseMovie.Celebrity{Name: director.Name, Resource: director.Resource, Gender: director.Gender})
		}
		for _, rank := range Movie.MovieRank {
			response.RankScores = append(response.RankScores, responseMovie.RankScore{Score: rank.Score, Organization: rank.Organization})
		}
		go func() {
			err := m.RedisRepository.SaveSpecificCache(id, response)
			if err != nil {
				serviceLogger := logger.ServiceLogger{Repository: m.RedisRepository}
				serviceLogger.Error(logUtils.LogData{Event: "Failure to save cache", User: "system", Details: err.Error()})
			}
		}()
		return response, nil
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, utils.ExistError{Message: "No such file with id"}
	} else {
		return response, err
	}
}

func (m *ServiceMovieSpecific) SetViewed(id string, user string) {
	HistoryMovie := model.UserHistoryMovie{MovieId: id, UserMail: user, CreatedAt: time.Now()}
	if err := m.SqlRepository.CreateOrUpdateUserHistoryMovie(HistoryMovie); err != nil {
		serviceLogger := logger.ServiceLogger{Repository: m.RedisRepository}
		serviceLogger.Error(logUtils.LogData{Event: "Failure to update history", Details: err.Error()})
	}
}
