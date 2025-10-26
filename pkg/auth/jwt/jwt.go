package jwt

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwtlib.RegisteredClaims
}

func Generate(subject, secret string, ttl time.Duration) (string, error) {
	now := time.Now()
	claims := Claims{RegisteredClaims: jwtlib.RegisteredClaims{
		Subject:   subject,
		IssuedAt:  jwtlib.NewNumericDate(now),
		ExpiresAt: jwtlib.NewNumericDate(now.Add(ttl)),
	}}
	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func Validate(tokenString, secret string) (*Claims, error) {
	parsed, err := jwtlib.ParseWithClaims(tokenString, &Claims{}, func(token *jwtlib.Token) (interface{}, error) {
		return []byte(secret), nil
	}, jwtlib.WithValidMethods([]string{jwtlib.SigningMethodHS256.Name}))
	if err != nil {
		return nil, err
	}
	if claims, ok := parsed.Claims.(*Claims); ok && parsed.Valid {
		return claims, nil
	}
	return nil, jwtlib.ErrTokenInvalidClaims
}
