package movie

import (
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func MiddleWareMovieIdContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := c.Param("movie_id")
		IdUuid, err := uuid.Parse(Id)

		if err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Parameter error not uuid"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("movie_id", IdUuid.String())
		c.Next()
	}
}
