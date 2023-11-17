package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DitoAdriel99/go-oauth2/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("err loading: %v", err)
	}
	handler := router.New()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")),
		Handler: handler,
	}

	// Start the server
	server.ListenAndServe()

}
