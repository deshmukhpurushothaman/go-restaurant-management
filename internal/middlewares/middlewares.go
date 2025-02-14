package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userContextKey contextKey = "user"

// Role-based constants
const (
	RoleAdmin = 0
	RoleUser  = 1
)

type BlacklistedToken struct {
	Token     string
	ExpiresAt time.Time
}

// Token blacklist stored in memory (you can use a Redis store for a production application)
var (
	Blacklist = make(map[string]BlacklistedToken) // Store token along with its expiration time
	Mu        sync.Mutex                          // Mutex to make the map concurrent-safe
)

const secretKey = "secret"

// AuthMiddleware validates the token and extracts claims.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Bearer token malformed", http.StatusUnauthorized)
			return
		}

		// Check if token is blacklisted
		Mu.Lock()
		if _, found := Blacklist[tokenString]; found {
			Mu.Unlock()
			http.Error(w, "Token has been revoked", http.StatusUnauthorized)
			return
		}
		Mu.Unlock()

		// Parse and validate token
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is correct
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Secret key used for signing
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the token is expired
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				http.Error(w, "Token has expired", http.StatusUnauthorized)
				return
			}
		} else {
			http.Error(w, "Token does not have an expiration claim", http.StatusUnauthorized)
			return
		}

		// Add claims to context for access in downstream handlers
		ctx := context.WithValue(r.Context(), userContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RoleMiddleware restricts access based on the user's role.
func RoleMiddleware(requiredRole uint) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get user claims from context
			claims, ok := r.Context().Value(userContextKey).(jwt.MapClaims)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Extract role from claims
			role, ok := claims["role"].(float64) // JWT numeric claims are float64
			if !ok || uint(role) != requiredRole {
				http.Error(w, "Forbidden: insufficient privileges", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func CleanupExpiredTokens() {
	for {
		time.Sleep(time.Hour * 24) // Clean up every day (adjust as needed)

		Mu.Lock()
		for token, data := range Blacklist {
			if time.Now().After(data.ExpiresAt) {
				delete(Blacklist, token) // Remove expired token
			}
		}
		Mu.Unlock()
	}
}
