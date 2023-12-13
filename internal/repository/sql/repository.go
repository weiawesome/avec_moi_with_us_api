package sql

import (
	"avec_moi_with_us_api/api/utils"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{db: utils.GetDB()}
}
