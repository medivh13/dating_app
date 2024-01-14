package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret_key")

// TokenClaims menyimpan klaim JWT
type TokenClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken membuat token JWT
func GenerateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &TokenClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
