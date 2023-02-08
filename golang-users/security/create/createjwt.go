package create

import (
	"golang-users/security/generatekeys"
	date "golang-users/utils/date_utils"
	"golang-users/utils/rest_errors"
	"log"

	"github.com/golang-jwt/jwt"
)

const (
	ExpiresAt = "ExpiresAt"
	USER_NAME = "Name"
	USER_ROLE = "Role"
)

type UserInfoInt interface {
	CreateUserInfo(string, Role, int) UserInfo
	getName() string
	getRole() string
	getHours() int
}

type UserInfo struct {
	name  string
	role  Role
	hours int
}

func (u UserInfo) getName() string {
	return u.name
}
func (u UserInfo) getRole() string {
	return u.role.String()
}
func (u UserInfo) getHours() int {
	return u.hours
}
func (u UserInfo) CreateUserInfo(name string, role Role, hours int) UserInfo {
	return UserInfo{name: name, role: role, hours: hours}
}

func CreateToken(user UserInfoInt) (*string, error) {
	if user == nil {
		return nil, rest_errors.NewBadRequestError("user info cannot be null")
	}

	t := jwt.New(jwt.SigningMethodRS256)
	t.Claims = jwt.MapClaims{
		ExpiresAt: date.AddHoursToUnixToString(user.getHours()),
		USER_NAME: user.getName(),
		USER_ROLE: user.getRole(),
	}
	log.Println("Signing token...")
	token, err := t.SignedString(generatekeys.PRIVATE_KEY)
	if err != nil {
		log.Println("Error signing token")
		return nil, err
	}
	log.Printf("Generated JWT: %v", token)
	return &token, err
}
