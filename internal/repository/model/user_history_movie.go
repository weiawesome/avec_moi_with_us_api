package model

import "time"

type UserHistoryMovie struct {
	UserMail  string `gorm:"primaryKey;size:254"`
	MovieId   string `gorm:"primaryKey;type:char(36)"`
	CreatedAt time.Time
}
