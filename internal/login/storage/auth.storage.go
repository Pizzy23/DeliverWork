package login

import (
	"deliver/models"

	"gorm.io/gorm"
)

func StoreToken(q *gorm.DB, email string, token string) error {
	user, err := FindUserByEmail(q, email)
	if err != nil {
		return err
	}
	auth := models.Auth{
		Token:  token,
		IdUser: user.IdUser,
	}
	if err := q.Create(&auth).Error; err != nil {
		return nil
	}
	return nil
}
