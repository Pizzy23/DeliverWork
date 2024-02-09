package login

import (
	"deliver/models"

	"gorm.io/gorm"
)

func CreateNewUser(q *gorm.DB, filters map[string]interface{}) (models.User, error) {
	newUser := models.User{
		Name:     filters["name"].(string),
		Email:    filters["email"].(string),
		Password: filters["password"].(string),
		IsLogged: false,
		IsValid:  false,
	}
	if err := q.Create(&newUser).Error; err != nil {
		return newUser, err
	}

	authToken := models.Auth{
		Token:  filters["token"].(string),
		IdUser: newUser.IdUser,
	}

	if err := q.Create(&authToken).Error; err != nil {
		return newUser, err
	}

	return newUser, nil
}

func ConnectUserInDB(q *gorm.DB, email string) (map[string]interface{}, error) {
	resultUser, errUser := FindUserByEmail(q, email)
	if errUser != nil {
		return nil, errUser
	}
	var login models.User
	if err := q.Model(&resultUser).Association("Login").Find(&login); err != nil {
		return nil, err
	}
	login.IsLogged = true
	if err := q.Save(&login).Error; err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"message": "Connected successfully",
	}

	return result, nil
}

func DesconnectUserInDB(q *gorm.DB, email string) (map[string]interface{}, error) {

	resultUser, errUser := FindUserByEmail(q, email)
	if errUser != nil {
		return nil, errUser
	}
	var login models.User
	if err := q.Model(&resultUser).Association("Login").Find(&login); err != nil {
		return nil, err
	}
	login.IsLogged = false
	if err := q.Save(&login).Error; err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"message": "Desconnected successfully",
	}

	return result, nil
}

func FindUserByEmail(q *gorm.DB, email string) (models.User, error) {
	var user models.User
	if err := q.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func SearchToken(q *gorm.DB, token string) (models.Auth, error) {
	var auth models.Auth
	if err := q.Where("token = ?", token).First(&auth).Error; err != nil {
		return auth, err
	}
	return auth, nil

}

func AuthUser(q *gorm.DB, email string, token string) error {
	resultUser, errUser := FindUserByEmail(q, email)
	if errUser != nil {
		return errUser
	}
	resultToken, errToken := SearchToken(q, token)

	if errToken != nil {
		return errToken
	}
	resultUser.IsLogged = true
	resultUser.IsValid = true
	if err := q.Save(&models.User{}).Error; err != nil {
		return err
	}

	if err := q.Delete(&resultToken).Error; err != nil {
		return err
	}
	return nil
}
