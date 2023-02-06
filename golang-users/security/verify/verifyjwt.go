package verify

import (
	"golang-users/security/generatekeys"
	resterrors "golang-users/utils/rest_errors"
	"log"

	"github.com/golang-jwt/jwt"
)

const (
	ExpiresAt = "ExpiresAt"
	USER_NAME = "Name"
	USER_ROLE = "Role"
)

func VerifyToken(tokenString string, r Role) (*string, resterrors.RestErr) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims,
		func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return generatekeys.PUBLIC_KEY, nil
		},
	)

	if err != nil {
		return nil, resterrors.NewUnauthorizedError("Token Signature is Not Valid") //UNAUTHORIZED 401
	}
	name := claims[USER_NAME].(string)
	role := claims[USER_ROLE].(string)

	log.Println("Token Claims[ Name: " + name + ", Role: " + role + " ]")

	if role != r.String() {
		return nil, resterrors.NewForbiddenError("Token Role Not Valid") //FORBIDDEN 403
	}

	return &name, nil // ok
}
