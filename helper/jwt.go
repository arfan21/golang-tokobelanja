package helper

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

func NewJwt(id uint, role string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	expired, err := strconv.Atoi(os.Getenv("JWT_EXPIRED"))

	if err != nil {
		return "", err
	}

	jwtExpired := time.Now().Local().Add(time.Minute * time.Duration(expired))

	claims := JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: jwtExpired},
			Issuer:    "hacktiv8-kanbanboard",
			Subject:   strconv.Itoa(int(id)),
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(secretKey)
}

func ParseJwt(tokenString string) (id uint, role string, err error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("invalid token")
	}

	idInt, err := strconv.Atoi(claims["sub"].(string))

	if err != nil {
		return 0, "", err
	}

	return uint(idInt), claims["role"].(string), nil
}
