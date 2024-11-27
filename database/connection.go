package database

import (
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func Connection() *sqlx.DB {
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	dsn := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	once.Do(func() {
		var err error
		db, err = sqlx.Connect("postgres", dsn)
		if err != nil {
			log.Fatal("Failed to connect to database", err)
		}
	})
	return db
}
