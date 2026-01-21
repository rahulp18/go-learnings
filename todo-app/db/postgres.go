package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgres() *sqlx.DB {
	dsn := "postgres://postgres:root@localhost:5432/todo?sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatal("Failed to connected to DB", err)
	}
	log.Println("Postgres connected")
	return db
}
