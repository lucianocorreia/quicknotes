package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucianocorreia/quicknotes/internal/handlers"
)

func main() {
	config := loadConfig()

	slog.SetDefault(newLogger(os.Stdout, config.GetLevelLog()))

	dbPool, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer dbPool.Close()

	// ping the database
	err = dbPool.Ping(context.Background())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Database connected")

	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("views/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	noteHandler := handlers.NewNoteHandler()

	mux.HandleFunc("/", noteHandler.NotesList)
	mux.Handle("/note/view", handlers.HandlerWithError(noteHandler.NotesView))
	mux.HandleFunc("/note/new", noteHandler.NotesNew)
	mux.HandleFunc("/note/create", noteHandler.NotesCreate)

	slog.Info(fmt.Sprintf("Starting server on port %s", config.ServerPort))
	err = http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux)
	if err != nil {
		panic(err)
	}
}
