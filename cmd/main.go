package main

import (
	"delivery-tracker-go/internal/models"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	cnf := loadConfig()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dbModel, err := models.InitDB(cnf.DB)
	if err != nil {
		slog.Error("Failed to initialize DB", "error", err)
		os.Exit(1)
	}
	slog.Error("DB initialized successfully", "error", err)

	RegisterCustomValidators()

	h := NewHandler(dbModel)

	router := gin.Default()
	if err := loadTemplates(router); err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}

	sessionStore := setupSessionStore(dbModel.DB, []byte(cnf.SessionSecretKey))
	setupRoutes(router, h, sessionStore)
	slog.Info("Server starting", "url", "http://localhost:"+cnf.Port)
	router.Run(":" + cnf.Port)

}
