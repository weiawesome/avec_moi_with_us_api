package model

type Celebrity struct {
	CelebrityId int32  `gorm:"primaryKey"`
	Name        string `gorm:"size:50;not null"`
	Gender      string `gorm:"type:enum('male','female','other');default:'other';not null"`
	Resource    string `gorm:"type:text"`

	Movies    []Movie `gorm:"many2many:actor;foreignKey:CelebrityId;joinForeignKey:CelebrityId;References:MovieId;JoinReferences:MovieId"`
	Directing []Movie `gorm:"many2many:director;foreignKey:CelebrityId;joinForeignKey:CelebrityId;References:MovieId;JoinReferences:MovieId"`
}
