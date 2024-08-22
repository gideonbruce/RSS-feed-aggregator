package main

import (
	"log"
	"net/http"
	"os"
     
	"github.com/go-chi/cors"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
 
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options {
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge:         300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}
     
	log.Printf("Server starting on port %v", port)
	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}

}