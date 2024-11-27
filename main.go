package main

import (
	"go_task/server"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := server.StartServer()
	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}
