package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	router.POST("mock/post", func(c *gin.Context) {
		body := make(map[string]string)
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		body["key1"] = body["key1"] + "_bye"
		body["key2"] = body["key2"] + "_bye"
		c.JSON(200, &body)
	})
	router.Run(":8081")
}
