package app

import (
	"golang-users/controllers/login"
	"golang-users/controllers/ping"
	"golang-users/controllers/users"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	router.GET("/ping", ping.Ping)
	router.GET("/rickandmorty", ping.RickAndMorty)
	router.POST("/test", ping.PostMock)
	router.GET("/login", login.Login)
	router.GET("/verifytoken", login.VerifyToken)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.GET("/users/gorm/email/:email", users.GetByEmail)
	router.POST("/users/gorm", users.CreateWithGorm)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)

	router.GET("/qr/:stringa", ping.QRcontroller)
}
