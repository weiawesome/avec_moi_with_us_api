package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type JwtClaim struct {
	Type  string `json:"type"`
	Fresh bool   `json:"fresh"`
	Csrf  string `json:"csrf"`
	jwt.RegisteredClaims
}

func GetJwt(mail string) (string, string, error) {
	jwtSecret := []byte(EnvJwtSecret())
	csrf := uuid.New().String()
	id := uuid.New().String()
	expiredDays, _ := strconv.Atoi(EnvJwtExpireDays())
	issuer := EnvJwtIssuer()

	claims := &JwtClaim{
		"access",
		false,
		csrf,
		jwt.RegisteredClaims{
			ID:        id,
			Subject:   mail,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiredDays) * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSecret)
	return ss, csrf, err
}

func ValidateJwt(jwtToken string) error {
	token, err := jwt.ParseWithClaims(jwtToken, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(EnvJwtSecret()), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(*JwtClaim); ok && token.Valid && claims.Issuer == EnvJwtIssuer() {
		return nil
	} else {
		return errors.New("content error")
	}
}

func ValidateJwtCsrf(jwtToken string, csrfToken string) error {
	token, err := jwt.ParseWithClaims(jwtToken, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) { return []byte(EnvJwtSecret()), nil })
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(*JwtClaim); ok {
		if token.Valid && claims.Issuer == EnvJwtIssuer() && claims.Csrf == csrfToken {
			return nil
		} else {
			return errors.New("jwt content error")
		}
	} else {
		return errors.New("jwt can't parse")
	}
}

func InformationJwt(jwtToken string) *JwtClaim {
	token, err := jwt.ParseWithClaims(jwtToken, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(EnvJwtSecret()), nil
	})
	if err != nil {
		return nil
	}
	if claims, ok := token.Claims.(*JwtClaim); ok && token.Valid && claims.Issuer == EnvJwtIssuer() {
		return claims
	} else {
		return nil
	}
}
