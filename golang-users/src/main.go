package main

import (
	_ "golang-users/docs" //NON OPZIONALE

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"golang-users/app"
	"golang-users/utils/logger"
)

var (
	router = gin.Default()
)

// swagger embed files

// @title           USERS API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes http
// @host      localhost:8080
// @BasePath  /
func main() {
	app.MapUrls(router)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	logger.Info("about to start the application...")
	router.Run(":8080")
}
