package security

import (
	"golang-users/security/generatekeys"
	"golang-users/utils/date_utils"
	"log"

	"github.com/golang-jwt/jwt"
)

type UserInfo struct {
	name  string
	role  Role
	hours int
}

type CustomClaims struct {
	*jwt.StandardClaims
	info UserInfo
}

func CreateToken(user UserInfo) (*string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = &CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: date_utils.AddHours(user.hours).Unix(),
		},
		user,
	}
	token, err := t.SignedString(generatekeys.SignKey)
	if err != nil {
		return nil, err
	}
	log.Printf("Generated JWT: %v", token)
	return &token, err
}
