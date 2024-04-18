package main

import (
	"fmt"
	"net/http"
	"os"

	"mainframe/account/config"
	"mainframe/account/handler"
)

func loadConfig(path string) {
	// Logging
	fmt.Println("Loading configuration")

	// Reading config file
	err := config.LoadConfig(path)
	if err != nil {
		fmt.Println("Error while loading configuration:", err)
		os.Exit(2)
	}

	// Logging
	fmt.Printf("Configuration ready: %+v\n", config.AppConfig)
}

func handleRequests() {
	// Logging
	fmt.Println("Setting up handlers")

	handler.HandleRequests()
}

func main() {
	// Logging
	fmt.Println("User service starting up")

	// Loading config file
	loadConfig("./config/config.yml")

	// Setting up routing paths
	handleRequests()

	// Starting up
	fmt.Println("Ready to listen incoming requests")
	http.ListenAndServe(":"+config.AppConfig.Server.Port, nil)
}
