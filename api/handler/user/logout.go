package user

import (
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"avec_moi_with_us_api/internal/service/logger"
	"avec_moi_with_us_api/internal/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerLogout struct {
	Service    user.ServiceLogout
	LogService logger.ServiceLogger
}

func (h *HandlerLogout) Handle(c *gin.Context) {
	jti, ok := c.Get("jti")
	if !ok {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if stringData, ok := jti.(string); ok {
		result, err := h.Service.Logout(stringData)
		if err == nil {
			go h.LogService.Info(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header})
			c.JSON(http.StatusNoContent, result)
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

}
