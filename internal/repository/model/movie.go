package model

type Movie struct {
	MovieId       string `gorm:"type:char(36);primaryKey"`
	Year          string `gorm:"type:char(4)"`
	Introduction  string `gorm:"type:text"`
	Resource      string `gorm:"type:text"`
	Score         float32
	Title         string `gorm:"size:255;not null"`
	OriginalTitle string `gorm:"size:255;not null"`

	MovieTitle   []MovieTitle `gorm:"foreignKey:MovieId;OnDelete: CASCADE"`
	Actors       []Celebrity  `gorm:"many2many:actor;foreignKey:MovieId;joinForeignKey:MovieId;References:CelebrityId;JoinReferences:CelebrityId;OnDelete: CASCADE"`
	Directors    []Celebrity  `gorm:"many2many:director;foreignKey:MovieId;joinForeignKey:MovieId;References:CelebrityId;JoinReferences:CelebrityId;OnDelete: CASCADE"`
	UsersLiked   []User       `gorm:"many2many:user_like_movies;foreignKey:MovieId;joinForeignKey:MovieId;References:Mail;JoinReferences:UserMail;AssociationForeignKey:Mail;AssociationJoinForeignKey:Mail;OnDelete: CASCADE"`
	UsersWatched []User       `gorm:"many2many:user_history_movies;foreignKey:MovieId;joinForeignKey:MovieId;References:Mail;JoinReferences:UserMail;AssociationForeignKey:Mail;AssociationJoinForeignKey:Mail;OnDelete: CASCADE"`
	MovieRank    []MovieRank  `gorm:"foreignKey:MovieId;OnDelete: CASCADE"`
	Genres       []Genre      `gorm:"many2many:movie_genres;foreignKey:MovieId;joinForeignKey:MovieId;References:GenreId;JoinReferences:GenreId;OnDelete: CASCADE"`
}
