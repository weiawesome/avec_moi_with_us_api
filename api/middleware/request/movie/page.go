package movie

import (
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func MiddleWarePagesContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		pageInt, err := strconv.Atoi(page)

		if err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Parameter error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("page", pageInt)
		c.Next()
	}
}
