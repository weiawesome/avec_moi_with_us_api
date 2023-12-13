package routes

import (
	handlerMovie "avec_moi_with_us_api/api/handler/movie"
	"avec_moi_with_us_api/api/middleware/jwt"
	requestMovie "avec_moi_with_us_api/api/middleware/request/movie"
	"avec_moi_with_us_api/internal/repository/chroma"
	"avec_moi_with_us_api/internal/repository/redis"
	"avec_moi_with_us_api/internal/repository/sql"
	"avec_moi_with_us_api/internal/service/logger"
	"avec_moi_with_us_api/internal/service/movie"
	"github.com/gin-gonic/gin"
)

func InitMovieRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository, chromaRepository *chroma.Repository) {
	middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	r.GET("", requestMovie.MiddleWarePagesContent(), requestMovie.MiddleWareRandomSeedContent(), (&handlerMovie.HandlerMovie{Service: movie.ServiceMovie{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.GET("/:movie_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestMovie.MiddleWareMovieIdContent(), (&handlerMovie.HandlerSpecificMovie{Service: movie.ServiceMovieSpecific{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.GET("/recently_hot", requestMovie.MiddleWarePagesContent(), (&handlerMovie.HandlerMovieRecentlyHot{Service: movie.ServiceMovieRecentlyHot{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.GET("/recently_view", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestMovie.MiddleWarePagesContent(), (&handlerMovie.HandlerMovieRecentlyView{Service: movie.ServiceMovieRecentlyView{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).View)
	r.GET("/recommend", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), (&handlerMovie.HandlerMovieRecommend{Service: movie.ServiceMovieRecommend{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.GET("/like", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestMovie.MiddleWarePagesContent(), (&handlerMovie.HandlerMovieLike{Service: movie.ServiceMovieLike{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).View)
	r.GET("/like/:movie_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestMovie.MiddleWareMovieIdContent(), (&handlerMovie.HandlerMovieLike{Service: movie.ServiceMovieLike{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Get)
	r.POST("/like/:movie_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestMovie.MiddleWareMovieIdContent(), (&handlerMovie.HandlerMovieLike{Service: movie.ServiceMovieLike{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Like)
	r.DELETE("/like/:movie_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestMovie.MiddleWareMovieIdContent(), (&handlerMovie.HandlerMovieLike{Service: movie.ServiceMovieLike{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).DisLike)
	r.GET("/search", requestMovie.MiddleWareContent(), requestMovie.MiddleWarePagesContent(), (&handlerMovie.HandlerMovieSearch{Service: movie.ServiceMovieSearch{SqlRepository: *sqlRepository, RedisRepository: *redisRepository, ChromaRepository: *chromaRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
}
