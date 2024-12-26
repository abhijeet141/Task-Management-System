package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	jwtsecret := os.Getenv("JWT_SECRET")
	if jwtsecret == "" {
		log.Fatal("Connection String is not set in the environment")
	}
	jwtSecret = []byte(jwtsecret)
}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Invalid Authorization Header", http.StatusUnauthorized)
			return
		}
		tokenId := strings.TrimPrefix(token, "Bearer ")
		tokenString, err := jwt.Parse(tokenId, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusForbidden)
			return
		}
		claims, ok := tokenString.Claims.(jwt.MapClaims)
		if ok && tokenString.Valid {
			username, found := claims["userName"].(string)
			if found {
				r.Header.Set("userName", username)
			}
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Invalid token claims", http.StatusForbidden)
	})
}
