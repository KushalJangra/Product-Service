package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB

func InitDB() {
	const (
		DBHost  = "127.0.0.1"
		DBUser  = "root"
		DBPass  = "Kush@123456"
		DBDbase = "pro"
	)

	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database ping error: %v", err)
	}

	Database = db
	log.Println("Database connected successfully!")
}
