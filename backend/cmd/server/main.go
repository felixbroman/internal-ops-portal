package main

import (
	"log"
	"net/http"

	"internal-ops-portal/internal/db"

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

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
