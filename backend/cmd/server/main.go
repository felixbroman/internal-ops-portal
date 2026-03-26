package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"internal-ops-portal/internal/auth"
	"internal-ops-portal/internal/db"
	"internal-ops-portal/internal/requests"
	"internal-ops-portal/internal/users"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	database := db.Connect()
	defer database.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// health
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// AUTH DOMAIN
	userRepo := users.NewPostgresRepository(database)
	authHandler := auth.NewHandler(userRepo)

	r.Route("/api/auth", func(authRouter chi.Router) {
		authRouter.Post("/signup", authHandler.Signup)
		authRouter.Post("/login", authHandler.Login)

		authRouter.Group(func(protected chi.Router) {
			protected.Use(auth.Middleware)
			protected.Get("/me", authHandler.Me)
		})
	})

	// REQUESTS DOMAIN
	requestRepo := requests.NewPostgresRepository(database)
	reqHandler := requests.NewHandler(requestRepo)

	r.Route("/api/requests", func(req chi.Router) {
		req.Group(func(protected chi.Router) {
			//employee
			protected.With(auth.RequireRole("employee")).
				Post("", reqHandler.Create)

			protected.With(auth.RequireRole("employee")).
				Get("/mine", reqHandler.Mine)

			// manager/admin
			protected.With(auth.RequireAnyRole("manager", "admin")).
				Get("", reqHandler.List)

			protected.With(auth.RequireAnyRole("manager", "admin")).
				Patch("/{id}", reqHandler.UpdateDecision)
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// frontend
	fs := http.FileServer(http.Dir("./frontend/dist"))
	r.Handle("/*", fs)

	log.Printf("Server running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
