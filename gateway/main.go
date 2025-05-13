package main

import (
	"log"
	"net/http"

	common "github.com/DavidReque/order-management-system/common"
	"github.com/joho/godotenv"
)

var httpAddr string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	httpAddr = common.EnvString("HTTP_ADDR", ":3000")

}

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Println("Starting server on", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
