package movie

import (
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TransformFilterValue(filters []string) ([]int, error) {
	var response []int
	for _, filter := range filters {
		if filterInt, err := strconv.Atoi(filter); err != nil {
			return response, err
		} else {
			response = append(response, filterInt)
		}
	}
	return response, nil
}
func MiddleWareFilterContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := c.QueryArray("filter")
		filtersInt, err := TransformFilterValue(filters)
		if err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Parameter error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		c.Set("filter", filtersInt)
		c.Next()
	}
}
