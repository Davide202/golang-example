package create

import (
	"golang-users/security/generatekeys"
	"golang-users/utils/date_utils"
	"log"

	"github.com/golang-jwt/jwt"
)

const (
	ExpiresAt = "ExpiresAt"
	USER_NAME = "Name"
	USER_ROLE = "Role"
)

type UserInfo struct {
	name  string
	role  Role
	hours int
}

func CreateUserInfo(name string, role Role, hours int) UserInfo {
	return UserInfo{name: name, role: role, hours: hours}
}

func CreateToken(user UserInfo) (*string, error) {
	t := jwt.New(jwt.SigningMethodRS256)
	t.Claims = jwt.MapClaims{
		ExpiresAt: date_utils.AddHours(user.hours).Unix(),
		USER_NAME: user.name,
		USER_ROLE: user.role,
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
