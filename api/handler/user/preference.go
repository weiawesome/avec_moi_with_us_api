package user

import (
	userRequest "avec_moi_with_us_api/api/request/user"
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"avec_moi_with_us_api/internal/service/logger"
	"avec_moi_with_us_api/internal/service/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerPreference struct {
	Service    user.ServicePreference
	LogService logger.ServiceLogger
}

func (h *HandlerPreference) Edit(c *gin.Context) {
	mail, okMail := c.Get("user")
	if !okMail {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	data, ok := c.Get("data")
	if !ok {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	if stringDataMail, ok := mail.(string); ok {
		if jsonData, ok := data.(userRequest.Preference); ok {
			result, err := h.Service.Edit(stringDataMail, jsonData)
			if err == nil {
				go h.LogService.Info(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataMail, Header: c.Request.Header})
				c.JSON(http.StatusOK, result)
			} else if errors.As(err, &responseUtils.ExistError{}) {
				go h.LogService.Warn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataMail, Header: c.Request.Header, Details: err.Error()})
				e := responseUtils.Error{Error: err.Error()}
				c.JSON(http.StatusUnauthorized, e)
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
func (h *HandlerPreference) Get(c *gin.Context) {
	mail, okMail := c.Get("user")
	if !okMail {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	if stringDataMail, ok := mail.(string); ok {
		result, err := h.Service.Get(stringDataMail)
		if err == nil {
			go h.LogService.Info(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataMail, Header: c.Request.Header})
			c.JSON(http.StatusOK, result)
		} else if errors.As(err, &responseUtils.ExistError{}) {
			go h.LogService.Warn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataMail, Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusUnauthorized, e)
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
}
func (h *HandlerPreference) GetType(c *gin.Context) {
	result, err := h.Service.GetType()
	if err == nil {
		go h.LogService.Info(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header})
		c.JSON(http.StatusOK, result)
	} else {
		go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
		e := responseUtils.Error{Error: err.Error()}
		c.JSON(http.StatusInternalServerError, e)
	}

}
