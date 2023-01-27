package users

import (
	"golang-users/datasources/mysql/gorm"
	"golang-users/domain/users"
	"golang-users/services"
	"golang-users/utils/logger"
	"golang-users/utils/rest_errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	logger.Info("Received " + user.String())
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	logger.Info("Saved " + result.String())
	c.JSON(http.StatusCreated, result)
}

func CreateWithGorm(c *gin.Context) {
	var user gorm.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	logger.Info("Received " + user.FirstName)
	saveErr := gorm.Create(&user)
	//saveErr := gorm.Update(&user)
	if saveErr != nil {
		c.JSON(500, saveErr)
		return
	}
	logger.Info("Saved " + user.FirstName)
	c.JSON(http.StatusCreated, user)
}

//		@Summary	GetUserById
//	 	@Param        user_id    path     string  true  "search by user_id"  Format(int)
//		@Schemes
//		@Description	description
//		@Tags			users
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	users.User
//		@Router			/users/{user_id} [get]
func Get(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetByEmail(c *gin.Context) {
	email := strings.TrimSpace(c.Param("email"))
	logger.Info("GetByEmail " + email)
	if email == "" {
		bdr := rest_errors.NewBadRequestError("user email should not be empty")
		c.JSON(bdr.Status(), bdr)
		return
	}
	users, err := gorm.GetUserByEmail(&email)
	if err != nil {
		ise := rest_errors.NewInternalServerError("error in get by email: "+email, err)
		c.JSON(ise.Status(), ise)
		return
	}
	c.JSON(200, users)
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, user)
}
