package Database

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func InitDB() {

	var err error

	DB, err = sql.Open("mysql","Suhail16:2002@tcp(127.0.0.1:3306)/jwt?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}else{
		log.Println("The database is connecting......")
	}

	err = DB.Ping()

	if err != nil{
		log.Fatal(err)
	}else{
		log.Println("The connection successfull")
	}
	
}