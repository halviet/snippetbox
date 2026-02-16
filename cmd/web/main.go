package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	l *slog.Logger
}

func main() {
	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	app := &application{l: logger}

	logger.Info(
		"starting server",
		slog.String("addr", *addr),
	)

	if err := http.ListenAndServe(*addr, app.routes()); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
