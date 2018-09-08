package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"impractical.co/auth/sessions"
)

func main() {
	key := os.Getenv("JWT_KEY")
	if key == "" {
		fmt.Println("JWT_KEY must be specified")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./spoof PROFILE_ID")
		os.Exit(1)
	}
	profile := os.Args[1]
	deps := sessions.Dependencies{
		JWTSecret: key,
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
