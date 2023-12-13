package routes

import (
	"avec_moi_with_us_api/internal/repository/chroma"
	"avec_moi_with_us_api/internal/repository/firebase"
	"avec_moi_with_us_api/internal/repository/redis"
	"avec_moi_with_us_api/internal/repository/sql"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}

	r.LoadHTMLGlob("templates/*")

	basicRouter := r.Group("/api/v1")

	userRouter := basicRouter.Group("/user")
	movieRouter := basicRouter.Group("/movie")
	apiDocsRouter := basicRouter.Group("/docs")

	sqlRepository := sql.NewRepository()
	redisRepository := redis.NewRepository()
	chromaRepository := chroma.NewRepository()
	firebaseRepository := firebase.NewRepository()

	InitUserRoutes(userRouter, sqlRepository, redisRepository, firebaseRepository)
	InitMovieRoutes(movieRouter, sqlRepository, redisRepository, chromaRepository)
	InitAPIDocsRoutes(apiDocsRouter)

	return r
}
