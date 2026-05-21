package main

import (
	"database/sql"
	"log"
	"project1/internal/database"
	"project1/internal/env"

	_ "modernc.org/sqlite"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database successfully.")

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-123456"),
		models:    models,
	}
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
