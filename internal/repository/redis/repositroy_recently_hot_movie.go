package redis

import (
	"avec_moi_with_us_api/api/response/movie"
	"avec_moi_with_us_api/internal/repository/utils"
	"context"
	"encoding/json"
)

func (r *Repository) GetRecentlyHot(page int) ([]movie.Movie, error) {
	redisKey := "RecentlyHot"
	var response []movie.Movie

	result, err := r.client.ZRangeWithScores(context.Background(), redisKey, int64(page), int64(page)).Result()
	if err != nil {
		return response, err
	}

	if len(result) > 0 {
		err := json.Unmarshal([]byte(result[0].Member.(string)), &response)
		if err != nil {
			return response, err
		}
		return response, nil
	} else {
		return response, utils.NotExistSource{Message: "The resource not exist."}
	}
}
