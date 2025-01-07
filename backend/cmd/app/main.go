package main

import (
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/arafetki/smartform.ai/backend/internal/app"
	"github.com/arafetki/smartform.ai/backend/internal/config"
	"github.com/arafetki/smartform.ai/backend/internal/db"
	"github.com/arafetki/smartform.ai/backend/internal/db/sqlc"
	"github.com/arafetki/smartform.ai/backend/internal/logging"
	"github.com/arafetki/smartform.ai/backend/internal/service"
)

func main() {

	cfg := config.Init()
	logger := logging.New(os.Stdout, slog.LevelInfo)

	// Set the log level based on the debug flag
	if cfg.App.Debug {
		logger.SetLevel(slog.LevelDebug)
	}

	// Connect to database
	db, err := db.Pool(cfg.Database.Dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	logger.Info("Database connection established sucessfully")

	svc := service.New(sqlc.New(db))

	app := app.New(cfg, logger, svc)

	if err := app.Run(); err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
