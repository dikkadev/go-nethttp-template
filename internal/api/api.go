package api

import (
	"CHANGEME/internal/middleware"
	"log/slog"
	"net/http"

	"github.com/sett17/prettyslog"
)

func Handler() http.Handler {
	apiRouter := http.NewServeMux()

	apiRouter.HandleFunc("/dbg", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Moin from API"))
	})

	logger := slog.New(prettyslog.NewPrettyslogHandler("API", prettyslog.WithLevel(slog.LevelDebug)))
	stack := middleware.CreateStack(
		middleware.Logger(logger, "api"),
	)

	return stack(apiRouter)
}
