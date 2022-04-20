package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// AuthMiddleware :This middleware handling authentication
func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get token from form value with key "auth"
		tokenString := r.FormValue("auth")
		if tokenString == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing access token")
			return
		}

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_TOKEN")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		handler.ServeHTTP(w, r)

	})
}
