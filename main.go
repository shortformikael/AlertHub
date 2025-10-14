package main

import (
	"fmt"
	"os"

	"github.com/shortformikael/AlertHub/src/config"
	"github.com/shortformikael/AlertHub/src/controllers/handler"
	"github.com/shortformikael/AlertHub/src/services/logger"
	"github.com/shortformikael/AlertHub/src/services/server"
)

var logComponent = "MAIN"

func main() {
	initialize()
	logger.Log(logComponent, logger.LogTypeInfo, "AlertHub starting...")

	handler.Start()
}

func initialize() {
	if err := config.Init(); err != nil {
		logger.Log(logComponent, logger.LogTypeFatal, fmt.Sprintf("Failed to initialize configuration: %v", err))
		os.Exit(1)
	}

	if err := server.Init(); err != nil {
		logger.Log(logComponent, logger.LogTypeFatal, fmt.Sprintf("Failed to initialize server: %v", err))
		os.Exit(1)
	}

	logger.Log(logComponent, logger.LogTypeInfo, "AlertHub configuration loaded successfully!")
}
