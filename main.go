package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"learningPlatform/api"
	"learningPlatform/config"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	router := mux.NewRouter()

	db, err := gorm.Open(postgres.Open(cfg.DBSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Disable table name pluralization
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	api.SetupRoutes(router, cfg, db)
	corsOpts := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Allow specific origin
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowCredentials(), // Allow credentials
	)

	log.Fatal(http.ListenAndServe(":8000", corsOpts(router)))
}
