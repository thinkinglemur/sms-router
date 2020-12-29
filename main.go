package main

import (
	"akon/bin/server"
	"akon/handlers/twilio"
	"fmt"
	"log"
	"net/http"
	"os"
)

const message = "This is a test"
const projectPrefix = "RoL"

func main() {

	logger := log.New(os.Stdout, projectPrefix, log.LstdFlags|log.Lshortfile)
	mux := http.NewServeMux()

	mux.HandleFunc("/", twilio.SmsWebhookHandler)

	srv := server.New(mux, ":8081")

	logger.Println("HTTP Server starting")
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
	fmt.Println(message)
}
