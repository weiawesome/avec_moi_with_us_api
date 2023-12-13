package routes

import (
	handlerUser "avec_moi_with_us_api/api/handler/user"
	"avec_moi_with_us_api/api/middleware/content_type"
	"avec_moi_with_us_api/api/middleware/jwt"
	requestUser "avec_moi_with_us_api/api/middleware/request/user"
	"avec_moi_with_us_api/internal/repository/firebase"
	"avec_moi_with_us_api/internal/repository/redis"
	"avec_moi_with_us_api/internal/repository/sql"
	"avec_moi_with_us_api/internal/service/logger"
	"avec_moi_with_us_api/internal/service/user"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository, firebaseRepository *firebase.Repository) {
	middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	r.GET("", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(),
		(&handlerUser.HandlerGetInfo{Service: user.ServiceGetInfo{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.POST("/signup", content_type.MiddleWareApplicationJson(), requestUser.MiddleWareSignupContent(),
		(&handlerUser.HandlerSignup{Service: user.ServiceSignup{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.POST("/login", content_type.MiddleWareApplicationJson(), requestUser.MiddleWareLoginContent(),
		(&handlerUser.HandlerLogin{Service: user.ServiceLogin{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.POST("/auth/firebase", content_type.MiddleWareApplicationJson(), requestUser.MiddleWareAuthFirebaseContent(),
		(&handlerUser.HandlerAuthFirebase{Service: user.ServiceAuthFirebase{FirebaseRepository: *firebaseRepository, SqlRepository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.DELETE("/logout", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(),
		(&handlerUser.HandlerLogout{Service: user.ServiceLogout{Repository: *redisRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.PUT("/edit_password", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareApplicationJson(), requestUser.MiddleWarePasswordEditContent(),
		(&handlerUser.HandlerEditPassword{Service: user.ServiceEditPassword{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.PUT("/edit_info", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareApplicationJson(), requestUser.MiddleWareInformationEditContent(),
		(&handlerUser.HandlerEditInfo{Service: user.ServiceEditInformation{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.GET("/preference/type",
		(&handlerUser.HandlerPreference{Service: user.ServicePreference{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).GetType)
	r.GET("/preference", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(),
		(&handlerUser.HandlerPreference{Service: user.ServicePreference{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Get)
	r.POST("/preference", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareApplicationJson(), requestUser.MiddleWarePreferenceContent(),
		(&handlerUser.HandlerPreference{Service: user.ServicePreference{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Edit)
}
