package main

import (
	"fmt"
	"os"

	"github.com/shortformikael/AlertHub/src/config"
	"github.com/shortformikael/AlertHub/src/controllers/handler"
	"github.com/shortformikael/AlertHub/src/services/logger"
)

var logComponent = "MAIN"

func main() {
	if err := config.Init(); err != nil {
		logger.Log(logComponent, logger.FATAL, fmt.Sprintf("Failed to initialize configuration: %v", err))
		os.Exit(1)
	}
	logger.Log(logComponent, logger.INFO, "AlertHub starting...")
	logger.Log(logComponent, logger.INFO, "AlertHub configuration loaded successfully!")

	handler.Start()
}
