package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Leo-Mart/goth-test/internal/middleware"
	"github.com/Leo-Mart/goth-test/internal/server"
	"github.com/Leo-Mart/goth-test/internal/store/db"
	"github.com/Leo-Mart/goth-test/internal/store/dbstore"
	"github.com/alexedwards/scs/v2"
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
	DB := db.Open()

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	mw, err := middleware.NewMiddleware(session)
	if err != nil {
		logger.Printf("could not get middleware")
		return
	}

	userDb := dbstore.NewUserStore(logger, DB)
	characterDb := dbstore.NewCharacterStore(logger, DB)

	srv, err := server.NewServer(logger, port, characterDb, userDb, session, mw)
	if err != nil {
		logger.Fatalf("Error when creating server: %s", err)
	}

	if err := srv.Start(); err != nil {
		logger.Fatalf("Error when starting server: %s", err)
		os.Exit(1)
	}
}
