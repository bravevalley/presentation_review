package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type Attendee struct {
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
	Name       string    `json:"name"`
	Location   string    `json:"location"`
	Review     string    `json:"review"`
}

var conn *pgx.Conn
var err error

func init() {
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}

func main() {
	defer conn.Close(context.Background())
	r := mux.NewRouter()

	r.HandleFunc("/" ,defaultHandler).Methods("GET")
	r.HandleFunc("/submit", submitHandler).Methods("POST")
	r.HandleFunc("/thx", thxHandler).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	log.Println("Server listening on :8080")
}
