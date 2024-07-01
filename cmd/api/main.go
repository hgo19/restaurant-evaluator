package main

import (
	"log"
	"net/http"
	"restaurant-evaluator/internal/domain/user"
	"restaurant-evaluator/internal/endpoints"
	"restaurant-evaluator/internal/infrastructure/database"
	"restaurant-evaluator/internal/infrastructure/utils"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userService := user.Service{
		Repository: &database.UserRepository{},
		Encrypter:  &utils.EncrypterBcrypt{},
	}

	handler := endpoints.Handler{
		UserService: userService,
	}

	r.Post("/signup", handler.UserPost)

	http.ListenAndServe(":3001", r)
}
