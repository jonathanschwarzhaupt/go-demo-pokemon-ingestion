package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/jonathanschwarzhaupt/go-demo-pokemon-ingestion/internal/database"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port   int
	env    string
	dbConn string
}

type application struct {
	config config
	logger *slog.Logger
	DB     *database.Queries
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "development", "Environment: (development|staging|production)")
	flag.StringVar(&cfg.dbConn, "dbConn", "", "Connection string to Turso database")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// initialize pool of database connections
	db, err := sql.Open("libsql", cfg.dbConn)
	if err != nil {
		log.Fatal("unable to open database connection: ", err)
	}
	dbQueries := database.New(db)

	app := &application{
		config: cfg,
		logger: logger,
		DB:     dbQueries,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
