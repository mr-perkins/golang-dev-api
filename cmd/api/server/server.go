package server

import (
	"fmt"
	"golang-dev-api/cmd/api/database"
	"net/http"
	"os"
	"strconv"
	"time"
)

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"time"

// 	"golang-dev-api/cmd/api/database"
// )

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("API_PORT"))
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", NewServer.port),
		// Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
