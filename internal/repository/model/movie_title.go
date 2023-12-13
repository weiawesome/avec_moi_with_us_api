package model

type MovieTitle struct {
	MovieTitleId string `gorm:"type:char(36);primaryKey"`
	MovieId      string `gorm:"size:254;foreignKey:MovieId;not null"`
	Title        string
	Language     string
}
