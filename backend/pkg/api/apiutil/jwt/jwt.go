package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"backend/pkg/user/role"
)

type Claims struct {
	jwt.StandardClaims
	UserId   string
	Username string
	Role     role.Role
}

var (
	secret     = []byte("JwTsEcReT")
	expireTime = 24 * time.Hour
	parser     = jwt.Parser{SkipClaimsValidation: true}
)

func GenerateToken(claims *Claims) string {
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(expireTime).Unix()
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	return token
}

func VerifyToken(token string) (*Claims, error) {
	if token == "" {
		return nil, errors.New("empty token")
	}

	claims := &Claims{}
	_, err := parser.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, errors.New("token is invalid")
	}

	if err := claims.Valid(); err != nil {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}
