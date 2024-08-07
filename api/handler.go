package api

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"learningPlatform/config"
	"learningPlatform/services"
	"net/http"
)

func SetupRoutes(router *mux.Router, cfg *config.Config, db *gorm.DB) {
	// Service
	userService := services.NewUserService(db, cfg)

	// Middlewares
	router.Use(func(next http.Handler) http.Handler {
		return JwtAuthentication(next, cfg.JWTSecretKey)
	})

	// Routes
	router.HandleFunc("/register", userService.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", userService.LoginHandler).Methods("POST")
	router.HandleFunc("/getMe", userService.GetMeHandler).Methods("GET", "POST")
}
