package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	config := serverConfig{
		PORT: os.Getenv("PORT"),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)

	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%s", config.PORT),
		Handler: mux,
	}

	log.Println(server.ListenAndServe())
	log.Println("Listening and serving at ", server.Addr)
}
