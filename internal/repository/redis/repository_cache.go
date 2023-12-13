package redis

import (
	"avec_moi_with_us_api/api/response/movie"
	"avec_moi_with_us_api/api/utils"
	"context"
	"encoding/json"
	"strconv"
	"time"
)

func (r *Repository) SaveMovieCache(page int, randomSeed int, movie movie.Movies) error {
	ctx := context.Background()
	key := "Movie-" + strconv.Itoa(page) + "-" + strconv.Itoa(randomSeed)
	bytes, err := json.Marshal(movie)
	if err != nil {
		return err
	}
	duration, err := strconv.Atoi(utils.EnvCacheDuration())
	if err != nil {
		duration = 5
	}
	expireDuration := time.Hour * time.Duration(duration)
	_, err = r.client.Set(ctx, key, string(bytes), expireDuration).Result()
	return err
}
func (r *Repository) SaveSearchCache(content string, page int, movie movie.Movies) error {
	ctx := context.Background()
	key := "Movie-" + content + "-" + strconv.Itoa(page)
	bytes, err := json.Marshal(movie)
	if err != nil {
		return err
	}
	duration, err := strconv.Atoi(utils.EnvCacheDuration())
	if err != nil {
		duration = 5
	}
	expireDuration := time.Hour * time.Duration(duration)
	_, err = r.client.Set(ctx, key, string(bytes), expireDuration).Result()
	return err
}
func (r *Repository) SaveSearchCandidateCache(content string, result []string) error {
	ctx := context.Background()
	key := "Movie-" + content
	bytes, err := json.Marshal(result)
	if err != nil {
		return err
	}
	duration, err := strconv.Atoi(utils.EnvCacheDuration())
	if err != nil {
		duration = 5
	}
	expireDuration := time.Hour * time.Duration(duration)
	_, err = r.client.Set(ctx, key, string(bytes), expireDuration).Result()
	return err
}
func (r *Repository) SaveSpecificCache(movieId string, specificMovie movie.SpecificMovie) error {
	ctx := context.Background()
	key := "Movie-" + movieId + "-Specific"
	bytes, err := json.Marshal(specificMovie)
	if err != nil {
		return err
	}
	duration, err := strconv.Atoi(utils.EnvCacheDuration())
	if err != nil {
		duration = 5
	}
	expireDuration := time.Hour * time.Duration(duration)
	_, err = r.client.Set(ctx, key, string(bytes), expireDuration).Result()
	return err
}

func (r *Repository) LoadMovieCache(page int, randomSeed int) (movie.Movies, error) {
	var v movie.Movies
	ctx := context.Background()
	key := "Movie-" + strconv.Itoa(page) + "-" + strconv.Itoa(randomSeed)
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return v, err
	}
	err = json.Unmarshal([]byte(result), &v)
	return v, err
}
func (r *Repository) LoadSearchCache(content string, page int) (movie.Movies, error) {
	var v movie.Movies
	ctx := context.Background()
	key := "Movie-" + content + "-" + strconv.Itoa(page)
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return v, err
	}
	err = json.Unmarshal([]byte(result), &v)
	return v, err
}
func (r *Repository) LoadSearchCandidateCache(content string) ([]string, error) {
	var v []string
	ctx := context.Background()
	key := "Movie-" + content
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return v, err
	}
	err = json.Unmarshal([]byte(result), &v)
	return v, err
}
func (r *Repository) LoadSpecificCache(movieId string) (movie.SpecificMovie, error) {
	var v movie.SpecificMovie
	ctx := context.Background()
	key := "Movie-" + movieId + "-Specific"
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return v, err
	}
	err = json.Unmarshal([]byte(result), &v)
	return v, err
}
