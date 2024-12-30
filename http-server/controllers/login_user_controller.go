package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	grpcclient "github/http-server/grpc-client"
	pb "github/http-server/proto/generated"
	"log"
	"net/http"
	"os"
	"time"

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

func GenerateJWT(userName string) (string, error) {
	claims := jwt.MapClaims{
		"username": userName,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("could not sign the token: %v", err)
	}
	return tokenString, nil
}
func GenerateRefresh(userName string) (string, error) {
	claims := jwt.MapClaims{
		"username": userName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshtoken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("could not sign the refresh token: %v", err)
	}
	return refreshtoken, nil
}
func LoginUserController(w http.ResponseWriter, r *http.Request) {
	var userInfo pb.UserInfo
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	userDetails := &pb.UserInfo{
		UserName: userInfo.UserName,
		Password: userInfo.Password,
	}
	client, err := grpcclient.TaskManagementClient()
	if err != nil {
		http.Error(w, "Failed to connect to gRPC service", http.StatusInternalServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.UserLogin(ctx, userDetails)
	if err != nil {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
		return
	}
	access_token, err := GenerateJWT(userInfo.UserName)
	if err != nil {
		http.Error(w, "Error generating access token", http.StatusInternalServerError)
		return
	}
	refresh_token, err := GenerateRefresh(userDetails.UserName)
	if err != nil {
		http.Error(w, "Error generating refresh token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"Access Token": access_token, "Refresh Token": refresh_token})
}
