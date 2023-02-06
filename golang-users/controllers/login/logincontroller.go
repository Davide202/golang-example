package login

import (
	create "golang-users/security/create"
	verify "golang-users/security/verify"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AUTHORIZATION = "Authorization"
	BEARER        = "Bearer "
)

func Login(c *gin.Context) {

	email := c.Query("email")
	password := c.Query("password")

	token, err := handleLogin(email, password)
	if err != nil {
		c.JSON(401, err.Error())
		return
	}
	c.Header(AUTHORIZATION, BEARER+*token)
	c.String(200, "http://localhost:8080/verifytoken")
}

func handleLogin(email, password string) (*string, error) {
	//RECUPERO LA PASSWORD, RUOLO E DELAY SUL DB E LA VERIFICO
	userRole := create.User
	//SE OK COSTRUISCO UserInfo E CREO IL TOKEN

	userInfo := create.CreateUserInfo(email, userRole, 3)

	token, err := create.CreateToken(userInfo)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func VerifyToken(c *gin.Context) {
	authorization := c.Request.Header.Get(AUTHORIZATION)
	tokenString := strings.SplitAfter(authorization, " ")[1]
	name, err := verify.VerifyToken(tokenString, verify.User)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.String(200, *name+" is authenticated")
}
