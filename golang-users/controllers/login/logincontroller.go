package login

import "github.com/gin-gonic/gin"

func login(c *gin.Context) {

	email := c.Query("email")
	password := c.Query("password")

	token, err := handleLogin(email, password)
	if err != nil {
		c.JSON(401, err.Error())
		return
	}
	c.Header("Authorization", "Bearer "+token)
	c.JSON(200, token)
}

func handleLogin(email, password string) (string, error) {
	//RECUPERO LA PASSWORD SUL DB E LA VERIFICO

	//SE OK

	return "Token", nil
}
