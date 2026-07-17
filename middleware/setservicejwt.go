package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var setKodeJWT = []byte("damarfinalproject123!")

type ClaimsDataObat struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, username string, role string) (string, error) {
	claims := ClaimsDataObat{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(setKodeJWT)
}

func ValidasiJWT(tokenString string) (*ClaimsDataObat, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ClaimsDataObat{}, func(t *jwt.Token) (interface{}, error) {
		return setKodeJWT, nil
	})

	if err != nil {
		return nil, err
	}

	cekClaims, ok := token.Claims.(*ClaimsDataObat)

	if !ok && !token.Valid {
		return nil, errors.New("token tidak valid atau sudah expired")
	}

	return cekClaims, nil
}
