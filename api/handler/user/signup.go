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

type HandlerSignup struct {
	Service    user.ServiceSignup
	LogService logger.ServiceLogger
}

func (h *HandlerSignup) Handle(c *gin.Context) {
	data, ok := c.Get("data")
	if !ok {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if jsonData, ok := data.(userRequest.SignUp); ok {
		result, err := h.Service.Signup(jsonData)
		if err == nil {
			go h.LogService.Info(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: jsonData.Mail, Header: c.Request.Header})
			c.JSON(http.StatusOK, result)
		} else if errors.As(err, &responseUtils.ExistError{}) {
			go h.LogService.Warn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: jsonData.Mail, Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusUnauthorized, e)
		} else {
			go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: jsonData.Mail, Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Data type mismatch"}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: jsonData.Mail, Header: c.Request.Header, Details: "Type Assertion error"})
		e := responseUtils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
