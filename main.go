package main

import (
	"github.com/shortformikael/AlertHub/src/config"
	"github.com/shortformikael/AlertHub/src/services/logger"
)

const logComponent = "MAIN"

func main() {
	config.Init()
	logger.Log(logComponent, logger.INFO, config.String())
	logger.Log(logComponent, logger.INFO, "AlertHub configuration loaded successfully!")

	logger.Log(logComponent, logger.INFO, "Testing colors")
	logger.Log(logComponent, logger.WARNING, "Testing colors")
	logger.Log(logComponent, logger.ERROR, "Testing colors")
	logger.Log(logComponent, logger.DEBUG, "Testing colors")
	logger.Log(logComponent, logger.FATAL, "Testing colors")
}
