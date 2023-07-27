package create

import (
	"jwt/model"
	"encoding/json"
	"log"
	"net/http"
	"jwt/database"

	_ "github.com/go-sql-driver/mysql"
)

//creat

func Createdomain(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var dom Model.LD
	_ = json.NewDecoder(r.Body).Decode(&dom)
	InsertdomainIntoDB(dom)
	json.NewEncoder(w).Encode(dom)


}
func InsertdomainIntoDB(domain Model.LD) {
	_, err := Database.DB.Exec("INSERT INTO jwt_auth (id, username, password) VALUES (?, ?, ?)",
	domain.ID, domain.USERNAME, domain.Password)
	if err != nil {
		log.Fatal(err)
	}
}
