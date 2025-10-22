package main

import (
	"fmt"
	com "mainframe-lib/common/config"
	"os"
	"processor/payment/config"
	"processor/payment/db"
	"processor/payment/service"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Get run args
	if len(os.Args) < 2 {
		fmt.Printf("No config file set\n")
		os.Exit(1)
	}
	configPath := os.Args[1]
	fmt.Printf("Loading configuration from %s\n", configPath)

	// Setup machine config
	var cfg config.Config
	err := com.LoadConfig(configPath, &cfg)
	if err != nil {
		fmt.Printf("Error while reading config file: %s\n", err.Error())
		os.Exit(2)
	}
	fmt.Printf("Configuration loaded: %+v\n", cfg)

	// Starting up
	fmt.Printf("Ready to listen incoming requests\n")

	// Main loop
	for {
		// Poll the queue for data
		key, paymentId, err := service.UnqueuePayment(cfg.Queue)
		if err != nil {
			fmt.Printf("Error while reading queue: %s\n", err.Error())
			continue
		}

		// Get abi value
		abi := strings.Split(key, ":")[0]
		if abi == "" || len(abi) != 5 {
			fmt.Printf("Queue content with key %s has invalid abi value\n", key)
			continue
		}

		fmt.Printf("Payment id %s received from abi %s\n", paymentId, abi)

		// Get payment data
		payment, err := db.SelectPayment(cfg.DB, abi, paymentId)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No payment with id %s found\n", paymentId)
			continue
		}
		if err != nil {
			fmt.Printf("Error while searching payment %s: %s\n", paymentId, err.Error())
			continue
		}

		// Process payment data
		outcome, err := service.ProcessPayment(cfg, abi, payment)
		if err != nil {
			fmt.Printf("Error while processing payment %s: %s\n", paymentId, err.Error())
			continue
		}

		// Update payment data
		payment.Outcome = outcome
		err = db.UpdatePayment(cfg.DB, abi, payment)
		if err != nil {
			fmt.Printf("Error while updating payment %s: %s\n", paymentId, err.Error())
			continue
		}

		fmt.Printf("Payment %s elaboration completed\n", paymentId)
	}
}
