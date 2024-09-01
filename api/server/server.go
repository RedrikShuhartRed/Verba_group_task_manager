package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
	"go.uber.org/zap"

	"github.com/RedrikShuhartRed/TaskManager/api/repository"
	"github.com/RedrikShuhartRed/TaskManager/api/routes"
	"github.com/RedrikShuhartRed/TaskManager/api/service"
	"github.com/RedrikShuhartRed/TaskManager/config"
	"github.com/RedrikShuhartRed/TaskManager/db"
)

// RunServer initializes and starts the HTTP server for the TaskManager application.
func RunServer() {

	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
	err := godotenv.Load(".env")
	if err != nil {
		zap.S().Errorf("error load .env: %v", err)
	}

	cfg := config.NewConfig()
	zap.S().Infof("Configuration loaded: %+v", cfg)

	port := cfg.Port

	storage, err := db.ConnectDB(cfg)
	if err != nil {
		zap.S().Errorf("Error connect DB, %v", err)
	}

	defer storage.CloseDB()

	r := gin.Default()
	taskRepo := repository.NewTaskRepository(storage.DB)
	taskService := service.NewTaskService(taskRepo)
	routes.RegisterRoutes(r, taskService)

	zap.S().Infof("Starting server on port %s...", port)

	err = r.Run(":" + port)
	if err != nil {
		zap.S().Fatal("Error starting Server, %v", err)
	}

}
