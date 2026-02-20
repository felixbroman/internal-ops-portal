package main

import (
	"log"
	"net/http"

	"internal-ops-portal/internal/auth"
	"internal-ops-portal/internal/db"
	"internal-ops-portal/internal/users"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	database := db.Connect()
	defer database.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	userRepo := users.NewPostpresRepository(database)
	authHandler := auth.NewHandler(userRepo)

	mux.HandleFunc("/api/auth/signup", authHandler.Signup)
	mux.HandleFunc("/api/auth/login", authHandler.Login)

	protected := auth.Middleware(http.HandlerFunc(authHandler.Me))
	mux.Handle("/api/auth/me", protected)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
