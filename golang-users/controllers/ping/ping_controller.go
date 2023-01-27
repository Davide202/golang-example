package ping

import (
	service "golang-users/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func RickAndMorty(c *gin.Context) {
	log.Printf("ClientIP: %s\n", c.ClientIP())
	//var params map[string]string
	//m := make(map[string]string)
	body, err := service.GetRickandmorty()
	if err != nil || body == nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, body)
	//c.Status(501)
}

func PostMock(c *gin.Context) {
	var array = []string{"", ""}
	er := c.ShouldBindJSON(&array)
	if er != nil {
		c.JSON(400, er.Error())
	}
	body, err := service.PostToMock(&array)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, body)
}
