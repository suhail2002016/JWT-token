package main

import (
	"fmt"
	"log"
	"net/http"

	// "encoding/json"
	"jwt/controller"
	"jwt/controller/handelers"
	"jwt/controller/creat"
	"jwt/database"
	// "jwt/model"

	"github.com/gorilla/mux"
) 

func main(){

	Database.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/fp_domaincheck", create.Createdomain).Methods("POST")

	r.HandleFunc("/login", Err.LoginHandler)	
	r.HandleFunc("/get-token", Jwt.GetJwtHandler)
	r.HandleFunc("/protected", Jwt.AuthMiddleware(Jwt.ProtectedHandler))
		
		fmt.Println("the server has started in http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	}



