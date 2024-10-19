package assets

import (
	"CHANGEME/internal/middleware"
	"embed"
	"log/slog"
	"net/http"

	"github.com/dikkadev/prettyslog"
)

//go:embed *
var Assets embed.FS

func Handler() http.Handler {
	logger := slog.New(prettyslog.NewPrettyslogHandler("ASSET"))
	stack := middleware.CreateStack(
		middleware.Logger(logger, "asset"),
		middleware.BlockPathEndingInSlash,
	)

	fs := http.FileServer(http.FS(Assets))
	return stack(http.StripPrefix("/assets/", fs))
}
