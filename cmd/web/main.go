package main

import (
	"database/sql"
	"flag"
	"github.com/halviet/snippetbox/internal/models"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	l *slog.Logger

	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":8080", "http service address")
	dsn := flag.String("dsn", "web:pass@tcp(127.0.0.1)/snippetbox?parseTime=true", "mysql DSN")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		l:        logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info(
		"starting server",
		slog.String("addr", *addr),
	)

	if err = http.ListenAndServe(*addr, app.routes()); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
