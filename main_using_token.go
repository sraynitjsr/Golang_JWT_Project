package main_using_token

import (
	"fmt"
	"net/http"
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

func makeAPIRequest(tokenString string) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+tokenString)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("API Response Status:", resp.Status)
}

func main_using_token() {
	tokenString, err := generateToken()
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}
	makeAPIRequest(tokenString)
}
