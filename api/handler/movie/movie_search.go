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

type HandlerMovieSearch struct {
	Service    movie.ServiceMovieSearch
	LogService logger.ServiceLogger
}

func (h *HandlerMovieSearch) Handle(c *gin.Context) {
	page, okPage := c.Get("page")
	if !okPage {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	content, okContent := c.Get("content")
	if !okContent {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	pageSize, err := strconv.Atoi(utils.EnvPageSize())
	if err != nil {
		pageSize = 10
	}
	if stringContent, ok := content.(string); ok {
		if intPage, ok := page.(int); ok {
			result, err := h.Service.Search(stringContent, intPage, pageSize)
			if err == nil {
				go h.LogService.Info(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header})
				c.JSON(http.StatusOK, result)
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
		go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Type Assertion error"})
		e := responseUtils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}

}
