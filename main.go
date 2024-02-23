package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	// secretKeyString := os.Getenv("SECRET_KEY")
	// fmt.Println(secretKeyString)
	portString := os.Getenv("PORT")
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Post("/tokens/{guid}", tokensHandler)
	v1Router.Post("/tokens/refresh", refreshTokensHandler)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Println("Server is running on port " + portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
