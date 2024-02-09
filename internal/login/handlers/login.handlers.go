package login

import (
	login "deliver/internal/login/services"

	"github.com/gin-gonic/gin"
)

// @Summary Conectar o usuario
// @Tags Login
// @Description Altera o modo de login para true
// @Accept json
// @Produce json
// @Param Email header string true "Email do usuario"
// @Param Password header string true "Senha do usuario"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/login [put]
func LoginUser(c *gin.Context) {
	email := c.GetHeader("email")
	login.Login(c, email)
}

// @Summary Desconectar o usuario
// @Tags Login
// @Description Altera o modo de login para false
// @Accept json
// @Produce json
// @Param Email header string true "Email do usuario"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/loggout [put]
func LoggedUser(c *gin.Context) {
	email := c.GetHeader("email")
	login.Logged(c, email)
}

// @Summary Criar token
// @Tags User
// @Description Criar o token e ser enviado ao usuario
// @Accept json
// @Produce json
// @Param Email header string true "Email do usuario"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/token [post]
func TokenUser(c *gin.Context) {
	email := c.GetHeader("email")
	login.GetAnotherToken(c, email)
}

// @Summary Autenticar User
// @Tags User
// @Description Autenticar o usuario atraves do email
// @Accept json
// @Produce json
// @Param Email header string true "Email do usuario"
// @Param Token header string true "Token do usuario"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/auth [put]
func AuthUser(c *gin.Context) {
	email := c.GetHeader("email")
	token := c.GetHeader("token")
	filters := map[string]interface{}{
		"email": email,
		"token": token,
	}
	login.Auth(c, filters)
}

// @Summary Criar um novo usuario
// @Tags User
// @Description Criar um usuario a partir das novas infos
// @Accept json
// @Produce json
// @Param Email header string true "Email do usuario"
// @Param Name header string true "Nome do usuario"
// @Param Password header string true "Senha do usuario"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/create-user [post]
func CreateUser(c *gin.Context) {
	name := c.GetHeader("Name")
	email := c.GetHeader("Email")
	password := c.GetHeader("Password")

	filters := map[string]interface{}{
		"name":     name,
		"email":    email,
		"password": password,
	}
	login.CreateUser(c, filters)
}

// @Summary Buscar um usuario
// @Tags User
// @Description Buscar um usuario
// @Accept json
// @Produce json
// @Param Email header string true "Email do usuario"
// @Success 200 {string} json "{"message": "Result"}"
// @Router /api/user [get]
func SeachUser(c *gin.Context) {
	email := c.GetHeader("Email")
	login.GetUser(c, email)
}
