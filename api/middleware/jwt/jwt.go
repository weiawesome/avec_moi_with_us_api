package jwt

import "avec_moi_with_us_api/internal/repository/redis"

type MiddlewareJwt struct {
	Repository *redis.Repository
}
