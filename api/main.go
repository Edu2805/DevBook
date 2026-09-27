package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	config.LoadConfigs()
	r := router.Generate()
	addr := fmt.Sprintf(":%d", config.Port)
	slog.Info("starting HTTP server", "address", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		slog.Error("HTTP server stopped", "error", err)
		os.Exit(1)
	}
}
