package main

import (
	"log"
	"os"

	"github.com/emmajiugo/social/internal/db"
	"github.com/emmajiugo/social/internal/env"
	"github.com/golang-migrate/migrate/v4"
	postgresDriver "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file FIRST
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// NOW get the environment variables after loading .env
	addr := env.GetString("DB_ADDR", "postgres://test:password@localhost:5432/social?sslmode=disable")
	maxOpenConns := env.GetInt("DB_MAX_OPEN_CONNS", 30)
	maxIdleConns := env.GetInt("DB_MAX_IDLE_CONNS", 30)
	maxIdleTime := env.GetString("DB_MAX_IDLE_TIME", "15m")

	db, err := db.New(addr, maxOpenConns, maxIdleConns, maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := postgresDriver.WithInstance(db, &postgresDriver.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration up completed successfully")
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration down completed successfully")
	}
}