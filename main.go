package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "tufan_ray"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte("my_secret_key"))
}

func main() {
	tokenString, err := generateToken()
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}
	fmt.Println("Generated Token:", tokenString)
}
