package main

import (
	"log"
	"os"

	"github.com/Leo-Mart/goth-test/internal/server"
	"github.com/Leo-Mart/goth-test/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	logger := log.New(os.Stdout, "[WoW Tracker]", log.LstdFlags)

	port := 8080

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Could not load env file")
	}
	logger.Print("Creating character store...")
	characterDb := store.NewCharacterStore(logger)

	srv, err := server.NewServer(logger, port, characterDb)
	if err != nil {
		logger.Fatalf("Error when creating server: %s", err)
	}

	if err := srv.Start(); err != nil {
		logger.Fatalf("Error when starting server: %s", err)
		os.Exit(1)
	}
}
