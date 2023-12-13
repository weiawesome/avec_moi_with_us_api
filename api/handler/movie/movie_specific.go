package movie

import (
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"avec_moi_with_us_api/internal/service/logger"
	"avec_moi_with_us_api/internal/service/movie"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerSpecificMovie struct {
	Service    movie.ServiceMovieSpecific
	LogService logger.ServiceLogger
}

func (h *HandlerSpecificMovie) Handle(c *gin.Context) {
	user, okUser := c.Get("user")
	if !okUser {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	id, okId := c.Get("movie_id")
	if !okId {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if stringDataUser, ok := user.(string); ok {
		if stringDataId, ok := id.(string); ok {
			result, err := h.Service.Get(stringDataId)
			if err == nil {
				go h.Service.SetViewed(stringDataId, stringDataUser)
				go h.LogService.Info(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header})
				c.JSON(http.StatusOK, result)
			} else if errors.As(err, &responseUtils.ExistError{}) {
				go h.LogService.Warn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
				e := responseUtils.Error{Error: err.Error()}
				c.JSON(http.StatusNotFound, e)
			} else {
				go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
				e := responseUtils.Error{Error: err.Error()}
				c.JSON(http.StatusInternalServerError, e)
			}
		} else {
			go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Type Assertion error"})
			e := responseUtils.Error{Error: "Internal error"}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: "Type Assertion error"})
		e := responseUtils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
