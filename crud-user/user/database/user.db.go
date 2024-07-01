package db

import (
	"database/sql"
	"log"
	"os"
	utils "crud-user/user/utils"
	_ "github.com/lib/pq"
)

// NewDB returns a new database connection.
func NewDB() *sql.DB {
	host := utils.GetEnv("DB_HOST", "127.0.0.1")
	port := utils.GetEnv("DB_PORT", "5432")
	user := utils.GetEnv("DB_USER", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "newpassword")
	dbname := utils.GetEnv("DB_NAME", "duonghdt")

	dbsource := "postgresql://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"

	db, err := sql.Open("postgres", dbsource)
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
		os.Exit(-1)
	}

	return db
}
