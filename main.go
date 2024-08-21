package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
 
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}
	
	router := chi.NewRouter()

	srv :=&http.Server{
		Handler: router,
		Addr:  ":"+ portString,
	}

	fmt.Println("PORT:", portString)
}