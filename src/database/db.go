package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

var db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load dotenv")
	}

	dbHost := os.Getenv("db_host")
	username := os.Getenv("db_user")
	dbName := os.Getenv("db_name")
	password := os.Getenv("db_pass")
	port := os.Getenv("db_port")

	// Creating connection
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, port)
	conn, err := sql.Open("postgres", dbUri)
	checkErr(err)

	db = conn

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	checkErr(err)

	// Applying migrations
	dir, err := os.Getwd()
	checkErr(err)

	m, err := migrate.NewWithDatabaseInstance(
		"file://" + dir + "/migrations",
		"postgres", driver)
	checkErr(err)

	dbVersion, err := strconv.ParseInt(os.Getenv("db_version"), 10, 32)
	checkErr(err)

	m.Steps(int(dbVersion))
}

func GetDB() *sql.DB {
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}