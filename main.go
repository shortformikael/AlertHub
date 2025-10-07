package main

import (
	"github.com/shortformikael/AlertHub/src/config"
	"github.com/shortformikael/AlertHub/src/services/logger"
)

const logComponent = "MAIN"

func main() {
	config.Init()
	logger.Log(logComponent, logger.INFO, "AlertHub starting...")
	logger.Log(logComponent, logger.INFO, "AlertHub configuration loaded successfully!")

}
