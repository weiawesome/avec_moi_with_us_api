package movie

import (
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"avec_moi_with_us_api/internal/service/logger"
	"avec_moi_with_us_api/internal/service/movie"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HandlerMovieRecentlyView struct {
	Service    movie.ServiceMovieRecentlyView
	LogService logger.ServiceLogger
}

func (h *HandlerMovieRecentlyView) View(c *gin.Context) {
	mail, okMail := c.Get("user")
	if !okMail {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	page, okPage := c.Get("page")
	if !okPage {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	pageSize, err := strconv.Atoi(utils.EnvPageSize())
	if err != nil {
		pageSize = 10
	}
	if stringDataMail, ok := mail.(string); ok {
		if intPage, ok := page.(int); ok {
			result, err := h.Service.Get(stringDataMail, intPage, pageSize)
			if err == nil {
				go h.LogService.Info(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataMail, Header: c.Request.Header})
				c.JSON(http.StatusOK, result)
			} else {
				go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataMail, Header: c.Request.Header, Details: err.Error()})
				e := responseUtils.Error{Error: err.Error()}
				c.JSON(http.StatusInternalServerError, e)
			}
		} else {
			go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataMail, Header: c.Request.Header, Details: "Type Assertion error"})
			e := responseUtils.Error{Error: "Internal error"}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataMail, Header: c.Request.Header, Details: "Type Assertion error"})
		e := responseUtils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
