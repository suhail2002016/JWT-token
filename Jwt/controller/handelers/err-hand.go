package Err

import (
	"encoding/json"
    "jwt/controller"
	// "jwt/model"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {


	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if Jwt.Authenticate(username, password) {
			// (username string, password string)
			tokenString, err := Jwt.GenerateToken(username)
			if err != nil {
				http.Error(w, "Failed to generate token", http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "Login successful!\n")
			fmt.Fprintf(w, "Token: %s", tokenString)
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	} else {
		http.ServeFile(w, r, "login.html")
	}
}

func SendResponse(w http.ResponseWriter, response interface{},code int) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, error interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(error)
}