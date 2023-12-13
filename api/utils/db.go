package utils

import (
	"avec_moi_with_us_api/internal/repository/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDB() (*gorm.DB, error) {
	user := EnvMySqlUser()
	password := EnvMySqlPassword()
	address := EnvMySqlAddress()
	dbName := EnvMySqlDb()

	dsn := user + ":" + password + "@tcp(" + address + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func InitDB() error {
	var err error
	if db, err = connectDB(); err != nil {
		return err
	}
	if err := db.AutoMigrate(&model.User{}, &model.RecommendMovie{}, &model.Movie{}, &model.MovieTitle{}, &model.Celebrity{}, &model.MovieRank{}, &model.Genre{}, &model.UserHistoryMovie{}, &model.UserLikeMovie{}); err != nil {
		return err
	}
	db = db.Debug()
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}
