package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	port = ":9090"
)

func main() {

	var db = getDb()


	handleMigration(db)
	sm := mux.NewRouter()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "hello")
	})

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:8081"}))

	// Create a new server
	s := http.Server{
		Addr:         ":9090",
		Handler:      ch(sm),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start the server
	go func() {
		log.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

func handleMigration(db *sql.DB) {
	shouldMigrate := flag.Bool("migrate", false, "Pass this flag to migrate")
	continueExecAfterMigrate := flag.Bool("continue", false, "Pass this flag to continue execution after migration")
	back := flag.Bool("rewind", false, "Pass this flag to rewind")
	steps := flag.Int("steps", 0, "Specify this number to decide how many steps to rewind")
	flag.Parse()

	if *shouldMigrate {
		println("Migrating")
		migrateDb(db, back, steps)
		if !*continueExecAfterMigrate {
			return
		}
	}
}

func getDb() *sql.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")

	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":3306)/"+database+"?parseTime=true&multiStatements=true")
	if err != nil {
		println("Could not connect to database")
		println(err.Error())
		log.Fatal(err)
	}
	var waited = 0
	for pingErr := db.Ping(); pingErr != nil; {
		println("Ping error: ", pingErr.Error())
		time.Sleep(1 * time.Second)
		waited++
		println("Waited ", waited, " seconds")
		if waited > 15 {
			log.Fatal("Could not get db connection")
		}
	}

	return db
}