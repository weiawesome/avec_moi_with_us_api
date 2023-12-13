package model

type MovieRank struct {
	MovieRankId         string `gorm:"type:char(36);primaryKey"`
	MovieId             string `gorm:"type:char(36);foreignKey:MovieId;not null"`
	OrganizationMovieId string
	Score               float32
	Organization        string
}
