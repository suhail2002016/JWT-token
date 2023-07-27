package Jwt

import (
	// "net/http"
	"database/sql"
	"fmt"
	Model "jwt/model"

	// "jwt/model"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, err := token.SignedString(Model.JwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetJwtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Access") == Model.JwtKey {
		username := r.FormValue("username")
		tokenString, err := GenerateToken(username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Token: %s", tokenString)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}




func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Token")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return Model.JwtSecretKey, nil
		})
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			username := claims["username"].(string)
			log.Printf("Authenticated user: %s", username)
			http.Error(w, "SUCCESS", http.StatusOK)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
	})
}



func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}





func Authenticate(username, password string) bool {
	
	query := "SELECT * FROM jwt_auth WHERE username=? AND password=?"
	row := Model.DB.QueryRow(query, username, password)

	var storedUsername, storedPassword string
	err := row.Scan(&storedUsername, &storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false // User not found
		}
		log.Fatal(err)
	}

	return true // Authentication successful
}


