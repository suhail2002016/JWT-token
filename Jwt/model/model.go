package Model

import (
	"database/sql"
	// "time"
)

type LD struct{
	ID          string         `json:"id"`
	USERNAME      string      `json:"username"`
	Password      string        `json:"password"`
 }

var DB *sql.DB
var JwtKey = "13579"
var JwtSecretKey = []byte(JwtKey)