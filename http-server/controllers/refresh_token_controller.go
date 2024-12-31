package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func RefreshTokenController(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.Header.Get("Authorization")

	if refreshToken == "" || !strings.HasPrefix(refreshToken, "Bearer ") {
		http.Error(w, "Invalid Authorization Header", http.StatusUnauthorized)
		return
	}

	tokenId := strings.TrimPrefix(refreshToken, "Bearer ")

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
		username, _ := claims["userName"].(string)
		accessToken, err := GenerateJWT(username)
		if err != nil {
			http.Error(w, "Error generating new access token", http.StatusInternalServerError)
			return
		}
		response, err := json.MarshalIndent(map[string]string{"Access-Token": accessToken}, " ", "\t")
		if err != nil {
			http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
