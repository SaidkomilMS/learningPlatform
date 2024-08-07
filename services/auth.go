package services

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"learningPlatform/config"
	"learningPlatform/models"
	"learningPlatform/repositories/implementations"
	"learningPlatform/repositories/interfaces"
	"net/http"
	"time"
)

type Claims struct {
	Username  string `json:"username"`
	IsStudent bool   `json:"is_student"`
	IsTeacher bool   `json:"is_teacher"`
	jwt.StandardClaims
}

type User struct {
	Username string `json:"username"`
}

type UserService struct {
	userRepo interfaces.UserRepository
	config   *config.Config
}

func NewUserService(db *gorm.DB, cfg *config.Config) *UserService {
	repo := implementations.NewGormUserRepository(db)
	return &UserService{
		userRepo: repo,
		config:   cfg,
	}
}

func (s *UserService) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var loginReq models.UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user := models.AuthUser{
		Username: loginReq.Username,
		Password: loginReq.Password, // You should hash this before saving
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Unable to encrypt password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	if err := s.userRepo.Create(ctx, &user); err != nil {
		http.Error(w, "Unable to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode("User successfully registered"); err != nil {
		http.Error(w, "Unable to return success result", http.StatusInternalServerError)
	}
}

func (s *UserService) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Authenticate the user
	user, err := s.userRepo.GetActiveUserByUsername(r.Context(), credentials.Username)
	if err != nil {
		http.Error(w, "Login failed", http.StatusUnauthorized)
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Login failed", http.StatusUnauthorized)
		return
	}

	if err := s.userRepo.UpdateLastLogin(r.Context(), user.ID); err != nil {
		http.Error(w, "Unable to update user data", http.StatusInternalServerError)
		return
	}

	// Create token
	expirationTime := time.Now().Add(time.Duration(s.config.JWTLifeTime)) // Token expiration time
	claims := &Claims{
		Username:  user.Username,
		IsStudent: user.IsStudent,
		IsTeacher: user.IsTeacher,
		StandardClaims: jwt.StandardClaims{
			Issuer:    s.config.JWTIssuer,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.JWTSecretKey))
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	// Return the token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func (s *UserService) GetMeHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the claims from context
	claims, ok := r.Context().Value("userInfo").(*Claims) // Cast to the appropriate type
	if !ok || claims == nil {
		http.Error(w, "Unauthorized or invalid token", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(claims)
}
