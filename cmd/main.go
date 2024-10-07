package main

import (
	"CHANGEME/internal/api"
	"CHANGEME/internal/middleware"
	"CHANGEME/web/assets"
	"log/slog"
	"net/http"

	"github.com/sett17/prettyslog"
)

func main() {
	slog.SetDefault(slog.New(prettyslog.NewPrettyslogHandler("APP", prettyslog.WithLevel(slog.LevelDebug))))
	slog.SetLogLoggerLevel(slog.LevelDebug)

	mainRouter := http.NewServeMux()

	httpRouter := http.NewServeMux()
	webLogger := slog.New(prettyslog.NewPrettyslogHandler("HTTP"))
	stack := middleware.CreateStack(
		middleware.Logger(webLogger, "req"),
	)

	mainRouter.Handle("/api/", http.StripPrefix("/api", api.Handler()))
	mainRouter.Handle("/assets/", assets.Handler())
	mainRouter.Handle("/", stack(httpRouter))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mainRouter,
	}

	slog.Info("Starting server on :8080")

	server.ListenAndServe()
}
