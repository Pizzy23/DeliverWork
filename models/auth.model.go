package models

type Auth struct {
	IdAuth uint   `gorm:"column:idAuth;primaryKey;autoIncrement" json:"id"`
	Token  string `gorm:"column:token;not null" json:"token"`
	IdUser uint   `gorm:"column:idUser" json:"IdUser"`
	User   User   `gorm:"foreignKey:IdUser" json:"user"`
}
