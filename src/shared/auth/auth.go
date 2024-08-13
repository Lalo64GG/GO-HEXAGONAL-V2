package auth

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(userID string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &jwt.StandardClaims{
        Subject:   userID,
        ExpiresAt: expirationTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
