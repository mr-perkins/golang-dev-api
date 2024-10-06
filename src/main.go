package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func connect() (*sql.DB, error) {
	bin, err := os.ReadFile("/run/secrets/db-password")
	if err != nil {
		return nil, err
	}
	return sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/dev-db-api?sslmode=disable", string(bin)))
}

func main() {
	log.Print("starting db")
	if err := prepare(); err != nil {
		log.Fatal(err)
	}

	slog.Info("starting server")

	http.HandleFunc("/", indexHandler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Printf("Error starting server: %s", err.Error())
		os.Exit(1)
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("request received")
	_, _ = fmt.Fprintf(w, "Request time: %s", time.Now().Format(time.RFC3339))
}

func prepare() error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	for i := 0; i < 60; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if _, err := db.Exec("DROP TABLE IF EXISTS blog"); err != nil {
		return err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS blog (id SERIAL, title VARCHAR)"); err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		if _, err := db.Exec("INSERT INTO blog (title) VALUES ($1);", fmt.Sprintf("Blog post #%d", i)); err != nil {
			return err
		}
	}
	return nil
}
