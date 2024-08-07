package api

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"learningPlatform/services"
	"net/http"
	"strings"
)

func JwtAuthentication(next http.Handler, jwtKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skipping middleware for no-auth routes
		noAuth := []string{"/login", "/register"} // List of endpoints that don't require auth
		requestPath := r.URL.Path                 // Current request path

		// Check if request does not need authentication
		for _, value := range noAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		// Proceed with JWT verification for other routes
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Expected header is "Bearer <token>"
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Invalid Authorization token format", http.StatusUnauthorized)
			return
		}

		tokenPart := bearerToken[1] // Extract the token part
		claims := &services.Claims{}

		token, err := jwt.ParseWithClaims(tokenPart, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil || !token.Valid { // Token is malformed or signature does not match
			http.Error(w, "Invalid Authorization token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userInfo", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
