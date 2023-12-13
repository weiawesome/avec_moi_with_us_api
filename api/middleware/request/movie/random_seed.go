package movie

import (
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RandomSeedValueCheck(RandomSeed int) bool {
	if RandomSeed < 0 || RandomSeed > 100 {
		return false
	} else {
		return true
	}
}
func MiddleWareRandomSeedContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		randomSeed := c.DefaultQuery("random_seed", "0")
		randomSeedInt, err := strconv.Atoi(randomSeed)

		if err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Parameter error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		if RandomSeedValueCheck(randomSeedInt) {
			c.Set("random_seed", randomSeedInt)
		} else {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Parameter error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		c.Next()

	}
}
