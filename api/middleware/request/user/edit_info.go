package user

import (
	"avec_moi_with_us_api/api/request/user"
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateInformationEdit(data user.EditInformation) error {
	if data.Name == "" {
		return errors.New("name can't be empty")
	} else if len(data.Name) > 30 {
		return errors.New("length of name can't large than 30 chars")
	} else if !(data.Gender == "male" || data.Gender == "female" || data.Gender == "other") {
		return errors.New("gender error type")
	}
	return nil
}

func MiddleWareInformationEditContent() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data user.EditInformation

		if err := c.ShouldBindJSON(&data); err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateInformationEdit(data)

		if err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("data", data)
		c.Next()
	}
}
