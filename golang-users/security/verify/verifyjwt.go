package security

import (
	"golang-users/security/generatekeys"
	"golang-users/utils/date_utils"
	resterrors "golang-users/utils/rest_errors"

	"github.com/golang-jwt/jwt"
)

/**/
type UserInfo struct {
	name  string
	role  Role
	hours int
}

type CustomClaims struct {
	*jwt.StandardClaims
	info UserInfo
}

func VerifyToken(tokenString string, role Role) resterrors.RestErr {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return generatekeys.VerifyKey, nil
		},
	)
	if err != nil {
		return resterrors.NewUnauthorizedError("Token Signature is Not Valid") //UNAUTHORIZED 401
	}
	claims := token.Claims.(*CustomClaims)
	isExpired := date_utils.IsExpired(claims.ExpiresAt)
	if isExpired {
		return resterrors.NewForbiddenError("Token is Expired") //FORBIDDEN 403
	}
	if claims.info.role == role {
		return nil // ok
	}
	return resterrors.NewForbiddenError("Token Role Not Valid") //FORBIDDEN 403
}
