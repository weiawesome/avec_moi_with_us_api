package model

type User struct {
	Mail     string `gorm:"primaryKey;size:254;"`
	Gender   string `gorm:"type:enum('male','female','other');default:'other';not null"`
	Salt     []byte `gorm:"type:Binary(16);default:null"`
	Password string
	Name     string `gorm:"size:50;not null"`

	RecommendMovie RecommendMovie
	LikeMovies     []Movie `gorm:"many2many:user_like_movies;foreignKey:Mail;joinForeignKey:UserMail;References:MovieId;JoinReferences:MovieId;AssociationForeignKey:MovieId;AssociationJoinForeignKey:MovieId"`
	HistoryMovies  []Movie `gorm:"many2many:user_history_movies;foreignKey:Mail;joinForeignKey:UserMail;References:MovieId;JoinReferences:MovieId;AssociationForeignKey:MovieId;AssociationJoinForeignKey:MovieId"`
	Genres         []Genre `gorm:"many2many:user_genres;foreignKey:Mail;joinForeignKey:UserMail;References:GenreId;JoinReferences:GenreId;AssociationForeignKey:GenreId;AssociationJoinForeignKey:GenreId"`
}
