package models

type User struct {
	IdUser   uint   `gorm:"column:idUser;primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"column:name;not null" json:"name"`
	Email    string `gorm:"column:email;not null" json:"email"`
	Password string `gorm:"column:password;not null" json:"password"`
	IsLogged bool   `gorm:"column:islogged;not null" json:"islogged"`
	IsValid  bool   `gorm:"column:isValid;not null" json:"isValid"`
}
