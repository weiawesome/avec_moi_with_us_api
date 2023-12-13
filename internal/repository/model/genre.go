package model

type Genre struct {
	GenreId uint   `gorm:"primaryKey"`
	Val     string `gorm:"not null"`

	Movies []Movie `gorm:"many2many:movie_genres;foreignKey:GenreId;joinForeignKey:GenreId;References:MovieId;JoinReferences:MovieId"`
	Users  []User  `gorm:"many2many:user_genres;foreignKey:GenreId;joinForeignKey:GenreId;References:Mail;JoinReferences:UserMail"`
}
