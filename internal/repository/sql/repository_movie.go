package sql

import (
	"avec_moi_with_us_api/internal/repository/model"
	StatusUtils "avec_moi_with_us_api/internal/repository/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math"
	"math/rand"
	"strings"
	"time"
)

func (r *Repository) ReadMovie(obj *model.Movie) error {
	if result := r.db.First(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return result.Error
	}
}
func (r *Repository) ReadLikeMovie(obj *model.UserLikeMovie) error {
	if result := r.db.First(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return result.Error
	}
}
func (r *Repository) PreLoadReadMovie(obj *model.Movie, preLoad []string) error {
	tx := r.db
	for _, p := range preLoad {
		tx = tx.Preload(p)
	}
	if result := tx.Find(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return result.Error
	}
}

func (r *Repository) ReadRecommendMovie(user string) ([]model.Movie, error) {
	u := model.User{Mail: user}
	r.db.Preload("RecommendMovie").First(&u)
	if _, err := uuid.Parse(u.RecommendMovie.RecommendId); err != nil {
		u.RecommendMovie.RecommendId = uuid.New().String()
	}
	if u.RecommendMovie.Result == nil || u.RecommendMovie.CreatedAt.Before(time.Now().Truncate(24*time.Hour)) {
		val, _, err := r.ReadMovieByPage(1, 10, rand.Intn(500)+100)
		if err != nil {
			return nil, err
		}
		marshal, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		u.RecommendMovie.CreatedAt = time.Now()
		u.RecommendMovie.Result = marshal

		r.db.Save(&u)
	}
	return u.RecommendMovie.FormatResult()

}
func (r *Repository) ReadLikeMovieByPage(user string, page int, pageSize int) ([]model.Movie, int, error) {
	var totalRecords int
	if err := r.db.Raw("SELECT COUNT(*) FROM user_like_movies WHERE user_mail = ?", user).Scan(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	var likedMovies []model.Movie
	if err := r.db.Raw("SELECT m.* FROM movies m INNER JOIN user_like_movies ulm ON m.movie_id = ulm.movie_id WHERE ulm.user_mail = ? ORDER BY ulm.created_at DESC LIMIT ? OFFSET ?", user, pageSize, (page-1)*pageSize).Scan(&likedMovies).Error; err != nil {
		return nil, 0, err
	}
	totalPages := totalRecords / pageSize
	if totalRecords%pageSize != 0 {
		totalPages++
	}

	return likedMovies, totalPages, nil
}
func (r *Repository) ReadRecentlyViewMovieByPage(user string, page int, pageSize int) ([]model.Movie, int, error) {
	var totalRecords int
	if err := r.db.Raw("SELECT COUNT(*) FROM user_history_movies WHERE user_mail = ?", user).Scan(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	var historyMovies []model.Movie
	if err := r.db.Raw("SELECT m.* FROM movies m INNER JOIN user_history_movies ulm ON m.movie_id = ulm.movie_id WHERE ulm.user_mail = ? ORDER BY ulm.created_at DESC LIMIT ? OFFSET ?", user, pageSize, (page-1)*pageSize).Scan(&historyMovies).Error; err != nil {
		return nil, 0, err
	}
	totalPages := totalRecords / pageSize
	if totalRecords%pageSize != 0 {
		totalPages++
	}

	return historyMovies, totalPages, nil
}
func (r *Repository) ReadMovieByPage(page int, pageSize int, randomSeed int) ([]model.Movie, int, error) {

	var movies []model.Movie
	var totalRecords int64
	if randomSeed == 0 {
		if err := r.db.Scopes(MovieModel(), OrderByScoreAndYear()).Count(&totalRecords).Scopes(Paginate(page, pageSize)).Find(&movies).Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err := r.db.Scopes(MovieModel(), OrderByRandom(randomSeed)).Count(&totalRecords).Scopes(Paginate(page, pageSize)).Find(&movies).Error; err != nil {
			return nil, 0, err
		}
	}

	totalPages := 0
	if totalRecords > 0 {
		totalPages = int(math.Ceil(float64(totalRecords) / float64(pageSize)))
	}

	return movies, totalPages, nil
}
func (r *Repository) ReadSearchMovieByPage(ids []string, title string, page int, pageSize int) ([]model.Movie, int, error) {
	var movies []model.Movie
	var totalRecords int64
	if len(ids) == 0 {
		if err := r.db.Scopes(MovieModel(), FilterByTitle(title), OrderByTitle(title)).Count(&totalRecords).Scopes(Paginate(page, pageSize)).Find(&movies).Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err := r.db.Scopes(MovieModel(), FilterByIdTitle(ids, title), OrderByIdTitle(ids, title)).Count(&totalRecords).Scopes(Paginate(page, pageSize)).Find(&movies).Error; err != nil {
			return nil, 0, err
		}
	}

	totalPages := 0
	if totalRecords > 0 {
		totalPages = int(math.Ceil(float64(totalRecords) / float64(pageSize)))
	}

	return movies, totalPages, nil
}

func (r *Repository) CreateOrUpdateUserLikeMovie(obj model.UserLikeMovie) error {
	var userLike model.UserLikeMovie
	r.db.FirstOrCreate(&userLike, obj)
	userLike.CreatedAt = time.Now()
	return r.db.Save(&userLike).Error
}
func (r *Repository) CreateOrUpdateUserHistoryMovie(obj model.UserHistoryMovie) error {
	var userHistory model.UserHistoryMovie
	r.db.FirstOrCreate(&userHistory, obj)
	userHistory.CreatedAt = time.Now()
	return r.db.Save(&userHistory).Error
}
func (r *Repository) DeleteMovieLike(obj model.UserLikeMovie) error {
	result := r.db.Delete(&obj)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return StatusUtils.NotExistSource{}
	}
	return nil
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
func OrderByScoreAndYear() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("score DESC, year DESC,movie_id ASC")
	}
}
func OrderByRandom(randomSeed int) func(db *gorm.DB) *gorm.DB {
	query := fmt.Sprintf("RAND(%d)", randomSeed)
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(query)
	}
}

func FilterByIdTitle(ids []string, title string) func(db *gorm.DB) *gorm.DB {
	titleLower := strings.ToLower(title)
	titleArg := "%" + titleLower + "%"
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("movie_id IN ?", ids).Or("LOWER(title) LIKE ?", titleArg).Or("LOWER(original_title) LIKE ?", titleArg)
	}
}
func OrderByIdTitle(ids []string, title string) func(db *gorm.DB) *gorm.DB {
	quoteIDs := func(ids []string) string {
		quoted := make([]string, len(ids))
		for i, id := range ids {
			quoted[i] = fmt.Sprintf("'%s'", id)
		}
		return strings.Join(quoted, ", ")
	}
	qId := quoteIDs(ids)
	titleLower := strings.ToLower(title)
	return func(db *gorm.DB) *gorm.DB {
		SQL_QUERY := fmt.Sprintf("CASE WHEN LOWER(title) LIKE '%%%s%%' OR LOWER(original_title) LIKE '%%%s%%' THEN 1 ELSE 2 END, FIELD(movie_id, %s), movie_id", titleLower, titleLower, qId)
		return db.Order(SQL_QUERY)
	}
}
func FilterByTitle(title string) func(db *gorm.DB) *gorm.DB {
	titleLower := strings.ToLower(title)
	titleArg := "%" + titleLower + "%"
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("LOWER(title) LIKE ?", titleArg).Or("LOWER(original_title) LIKE ?", titleArg)
	}
}
func OrderByTitle(title string) func(db *gorm.DB) *gorm.DB {
	titleLower := strings.ToLower(title)
	return func(db *gorm.DB) *gorm.DB {
		SQL_QUERY := fmt.Sprintf("CASE WHEN LOWER(title) = '%%%s%%' OR LOWER(original_title) = '%%%s%%'", titleLower, titleLower)
		return db.Order(SQL_QUERY)
	}
}
func MovieModel() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Model(&model.Movie{})
	}
}
