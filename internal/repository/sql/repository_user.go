package sql

import (
	"avec_moi_with_us_api/internal/repository/model"
	StatusUtils "avec_moi_with_us_api/internal/repository/utils"
	"errors"
	"gorm.io/gorm"
)

func (r *Repository) CreateUser(obj model.User) error {
	return r.db.Create(obj).Error
}
func (r *Repository) ReadUser(obj *model.User) error {
	if result := r.db.First(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return result.Error
	}
}
func (r *Repository) UpdateUser(obj model.User) error {
	return r.db.Model(&obj).Updates(obj).Error
}
func (r *Repository) DeleteUser(obj model.User) error {
	return r.db.Delete(&obj).Error
}

func (r *Repository) PreLoadReadUser(obj *model.User, preLoad []string) error {
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

func (r *Repository) UpdateUserGenres(mail string, newGenreIDs []uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var user model.User
		if err := tx.Where("mail = ?", mail).Preload("Genres").First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return StatusUtils.NotExistSource{Message: "Resource is not exist"}
		}
		if len(user.Genres) > 0 {
			if err := tx.Model(&user).Association("Genres").Clear(); err != nil {
				return err
			}
		}
		if len(newGenreIDs) > 0 {
			var genresToAdd []model.Genre
			if err := tx.Find(&genresToAdd, newGenreIDs).Error; err != nil {
				return err
			}
			if err := tx.Model(&model.User{Mail: mail}).Association("Genres").Append(genresToAdd); err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *Repository) GetGenres(obj *[]model.Genre) error {
	return r.db.Find(obj).Error
}
