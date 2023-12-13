package jwt

import (
	responseUtils "avec_moi_with_us_api/api/response/utils"
	"avec_moi_with_us_api/api/utils"
	"avec_moi_with_us_api/internal/repository/redis"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func validateBlackList(jwtToken string, repository *redis.Repository) error {
	jwt := utils.InformationJwt(jwtToken)
	if jwt == nil {
		return errors.New("JwtToken error")
	}
	jti := jwt.ID
	return repository.ValidateInBlacklist(jti)
}
func validateContent(jwtToken string) error {
	return utils.ValidateJwt(jwtToken)
}
func validateCookie(jwtToken string, csrfToken string) error {
	return utils.ValidateJwtCsrf(jwtToken, csrfToken)
}
func (m *MiddlewareJwt) JwtSecure() gin.HandlerFunc {
	return func(c *gin.Context) {
		statusContent := true

		jwtTokenContent := c.GetHeader("Authorization")
		if strings.HasPrefix(jwtTokenContent, "Bearer ") {
			jwtTokenContent, _ = strings.CutPrefix(jwtTokenContent, "Bearer ")
			statusContent = statusContent && (validateBlackList(jwtTokenContent, m.Repository) == nil) && (validateContent(jwtTokenContent) == nil)
		} else {
			statusContent = false
		}

		if statusContent {
			c.Set("jwt", jwtTokenContent)
			c.Next()
		} else {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Jwt validation fail"})
			e := responseUtils.Error{Error: "Invalidate jwtToken"}
			c.JSON(http.StatusUnprocessableEntity, e)
			c.Abort()
			return
		}

	}
}
