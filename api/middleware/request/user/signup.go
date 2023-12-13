package user

import (
	"avec_moi_with_us_api/api/request/user"
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/mail"
	"strings"
)

func validateSignup(data user.SignUp) error {
	if data.Mail == "" {
		return errors.New("mail can't be empty")
	} else if data.Password == "" {
		return errors.New("password can't be empty")
	} else if len(data.Password) < 8 {
		return errors.New("length of password can't shorten than 8 chars")
	} else if len(data.Password) > 30 {
		return errors.New("length of password can't larger than 30 chars")
	} else if strings.ContainsAny(data.Password, " ") {
		return errors.New("password can't contain space")
	} else if _, err := mail.ParseAddress(data.Mail); err != nil {
		return errors.New("mail can't parse")
	} else if !(data.Gender == "male" || data.Gender == "female" || data.Gender == "other") {
		return errors.New("gender error type")
	} else if data.Name == "" {
		return errors.New("name can't be empty")
	} else if len(data.Name) > 30 {
		return errors.New("length of name can't large than 30 chars")
	}
	return nil
}

func MiddleWareSignupContent() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data user.SignUp
		if err := c.ShouldBindJSON(&data); err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateSignup(data)

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
