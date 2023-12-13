package movie

import (
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWareContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		content := c.DefaultQuery("content", "")

		if content == "" {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Search empty!"})
			e := responseUtils.Error{Error: "Parameter error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("content", content)
		c.Next()
	}
}
