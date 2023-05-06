package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// print token
	fmt.Println("Token just after creating =>", token)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "tufan_ray"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// print token
	fmt.Println("Token after adding claim =>", token)

	// sign token with secret key
	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		fmt.Println("Error generating token string:", err)
		return
	}

	// print token
	fmt.Println("Token after adding claim and signing with secret =>", tokenString)
}
