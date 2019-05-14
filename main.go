package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"impractical.co/auth/sessions"
)

func pathOrContents(in string) (string, error) {
	if _, err := os.Stat(in); err == nil {
		contents, err := ioutil.ReadFile(in)
		if err != nil {
			return string(contents), err
		}
		return string(contents), nil
	}
	return in, nil
}

func main() {
	key := os.Getenv("JWT_KEY")
	if key == "" {
		fmt.Println("JWT_KEY must be specified")
		os.Exit(1)
	}
	privateKeyStr, err := pathOrContents(key)
	if err != nil {
		fmt.Println("Error getting JWT_KEY contents:", err)
		os.Exit(1)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyStr))
	if err != nil {
		fmt.Println("Error parsing JWT_KEY contents:", err)
		os.Exit(1)
	}
	serviceID := os.Getenv("SERVICE")
	if serviceID == "" {
		fmt.Println("SERVICE must be specified")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./spoof PROFILE_ID")
		os.Exit(1)
	}
	profile := os.Args[1]
	deps := sessions.Dependencies{
		JWTPrivateKey: privateKey,
		JWTPublicKey:  privateKey.Public().(*rsa.PublicKey),
		ServiceID:     serviceID,
	}
	token := sessions.AccessToken{
		ID:          "spoof",
		CreatedFrom: "spoof",
		ProfileID:   profile,
		ClientID:    "spoof",
		CreatedAt:   time.Now(),
	}
	jwt, err := deps.CreateJWT(context.Background(), token)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(jwt)
}
