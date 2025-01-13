package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

var (
	port = os.Getenv("API_PORT")
)

func gracefulShutdown(apiServer *http.Server) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
}

func main() {
	slog.Info("starting server")

	http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Printf("Error starting server: %s", err.Error())
		os.Exit(1)
	}

	// server := server.NewServer()

	// go gracefulShutdown(server)

	// err := server.ListenAndServe()
	// if err2 != nil && err != http.ErrServerClosed {
	// 	panic(fmt.Sprintf("http server error: %s", err))
	// }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("request received")
	_, _ = fmt.Fprintf(w, "Request time: %s", time.Now().Format(time.RFC3339))
}

// func connect() (*sql.DB, error) {
// 	// bin, err := os.ReadFile("/run/secrets/db-password")
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// return sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/dev-db-api?sslmode=disable", string(bin)))
// 	return sql.Open("postgres", "postgres://postgres:postgres@database:5432/dev-db?sslmode=disable")
// }

// func main() {
// 	log.Print("starting db")
// 	if err := prepare(); err != nil {
// 		log.Fatal(err)
// 	}

// 	slog.Info("starting server")

// 	http.HandleFunc("/", indexHandler)

// 	if err := http.ListenAndServe(":8000", nil); err != nil {
// 		log.Printf("Error starting server: %s", err.Error())
// 		os.Exit(1)
// 	}

// }

// func prepare() error {
// 	db, err := connect()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	for i := 0; i < 60; i++ {
// 		if err := db.Ping(); err == nil {
// 			break
// 		}
// 		time.Sleep(time.Second)
// 	}

// 	if _, err := db.Exec("DROP TABLE IF EXISTS blog"); err != nil {
// 		return err
// 	}

// 	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS blog (id SERIAL, title VARCHAR)"); err != nil {
// 		return err
// 	}

// 	for i := 0; i < 5; i++ {
// 		if _, err := db.Exec("INSERT INTO blog (title) VALUES ($1);", fmt.Sprintf("Blog post #%d", i)); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
