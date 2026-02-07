package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

/*
	type apiConfig struct {
		DB *database.Queries
	}
*/
func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Printf("Running Web Server on Port: %v", port)

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	/*
		conn, dbErr := sql.Open("postgres", dbUrl)

		if dbErr != nil {
			log.Fatal("Can't connect to Database")
		}

		apiCfg := apiConfig{
			DB: database.New(conn),
		}
	*/
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
	v1Router.Get("/health", handlerReadiness)
	//v1Router.Post("/users/create", apiCfg.handlerCreateUsers)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	srvErr := server.ListenAndServe()
	if srvErr != nil {
		log.Fatal(srvErr)
	}
}
