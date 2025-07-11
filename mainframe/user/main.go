package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"mainframe/user/config"
	"mainframe/user/handler"
)

func main() {
	// Get run args
	if len(os.Args) < 2 {
		fmt.Printf("No config file set\n")
		os.Exit(1)
	}
	configPath := os.Args[1]

	// Setup machine config
	fmt.Printf("Loading configuration from %s\n", configPath)
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Error while reading config file: %s\n", err.Error())
		os.Exit(2)
	}
	fmt.Printf("Configuration loaded: %+v\n", cfg)

	// Create the mux
	mux := http.NewServeMux()

	// Setup server routes
	fmt.Printf("Setting up routes\n")
	handler.SetupRoutes(cfg, mux)

	// Create the server
	server := &http.Server{
		Handler: mux,
	}
	ln, err := net.Listen("tcp", ":"+cfg.Server.Port)
	if err != nil {
		fmt.Printf("Error while assigning server port: %s\n", err.Error())
		os.Exit(3)
	}

	// Starting up
	fmt.Printf("Ready to listen incoming requests\n")
	server.Serve(ln)
	if err != nil {
		fmt.Printf("Error while starting up server: %s\n", err.Error())
		os.Exit(4)
	}
}
