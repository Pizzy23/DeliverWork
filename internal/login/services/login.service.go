package login

import (
	"deliver/db"
	login "deliver/internal/login/storage"

	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context, email string) {
	result, err := login.ConnectUserInDB(db.Repo, email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result})
}

func Logged(c *gin.Context, email string) {
	result, err := login.DesconnectUserInDB(db.Repo, email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result})
}

func CreateUser(c *gin.Context, filters map[string]interface{}) {
	password, ok := filters["password"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be a string"})
		return
	}

	filters["token"] = CreateToken()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Internal Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	filters["password"] = hashedPassword
	result, err := login.CreateNewUser(db.Repo, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result, "Response": "User created"})
}

func Auth(c *gin.Context, filters map[string]interface{}) {
	token := filters["token"].(string)
	result, err := login.SearchToken(db.Repo, token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.Token != token {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Token is not valid"})
		return
	}

	resultAuth := login.AuthUser(db.Repo, filters["email"].(string), filters["token"].(string))
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": resultAuth})
		return
	}

}

func CreateToken() string {
	Token := strconv.Itoa(rand.Intn(900000) + 100000)

	result, err := login.SearchToken(db.Repo, Token)
	if err == nil {
		return Token
	}

	if result.Token == Token {
		CreateToken()
	}

	return Token
}

func GetAnotherToken(c *gin.Context, email string) {
	token := CreateToken()
	err := login.StoreToken(db.Repo, email, token)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "New token send for your email"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func GetUser(c *gin.Context, email string) {
	res, err := login.FindUserByEmail(db.Repo, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": res})
}
