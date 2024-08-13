package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        tokenString := authHeader[len("Bearer "):]

        claims := &jwt.StandardClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        c.Set("userID", claims.Subject)
        c.Next()
    }
}
