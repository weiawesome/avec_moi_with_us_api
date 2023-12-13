package content_type

import (
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWareApplicationJson() gin.HandlerFunc {
	return func(c *gin.Context) {

		contentType := c.ContentType()

		if contentType != "application/json" {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Content-Type not application/json"})
			e := responseUtils.Error{Error: "Content-Type must be application/json not " + contentType}
			c.JSON(http.StatusUnsupportedMediaType, e)
			c.Abort()
			return
		}

		c.Next()
	}
}
