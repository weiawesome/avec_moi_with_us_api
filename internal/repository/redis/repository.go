package redis

import (
	"avec_moi_with_us_api/api/utils"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	client *redis.Client
}

func NewRepository() *Repository {
	return &Repository{client: utils.GetRedisClient()}
}
