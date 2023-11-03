package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTWrapper struct {
	SecretKey         string
	Issuer            string
	ExpirationMinutes int64
	ExiprationHours   int64
}

type JWTClaims struct {
	Email string
	jwt.StandardClaims
}

func (j *JWTWrapper) GenrateJwtToken(email string) (string, error) {
	claims := &JWTClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(j.ExpirationMinutes)).Unix(),
			Issuer:    email,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return signedToken, err
	}
	return signedToken, err
}

func (j *JWTWrapper) ValidateToken(signedToken string) (claims *JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		err := errors.New("Couldn't parse claims")
		return claims, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err := errors.New("Token expired")
		return claims, err
	}

	return

}
