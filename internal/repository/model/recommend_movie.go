package model

import (
	"encoding/json"
	"time"
)

type RecommendMovie struct {
	RecommendId string    `gorm:"primaryKey;type:char(36);not null"`
	UserMail    string    `gorm:"size:254;foreignKey:Mail;not null"`
	CreatedAt   time.Time `gorm:"not null"`
	Result      []byte    `gorm:"type:json;default:null"`
	Status      bool      `gorm:"not null"`
}

func (r RecommendMovie) FormatResult() ([]Movie, error) {
	var result []Movie
	err := json.Unmarshal(r.Result, &result)
	return result, err
}
