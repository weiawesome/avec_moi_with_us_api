package jwt

import (
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *MiddlewareJwt) JwtInformation() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, ok := c.Get("jwt")
		if !ok {
			go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
			e := responseUtils.Error{Error: "Data not found in context"}
			c.JSON(http.StatusInternalServerError, e)
			return
		}
		if stringData, ok := jwt.(string); ok {
			jwtToken := utils.InformationJwt(stringData)
			if jwtToken == nil {
				go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Jwt is empty"})
				e := responseUtils.Error{Error: "Jwt can't parse"}
				c.JSON(http.StatusUnprocessableEntity, e)
				return
			} else {
				c.Set("user", jwtToken.Subject)
				c.Set("jti", jwtToken.ID)
				c.Next()
			}
		} else {
			go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Type Assertion error"})
			e := responseUtils.Error{Error: "Data can't parse"}
			c.JSON(http.StatusInternalServerError, e)
			return
		}
	}
}
