package handler

import (
	"github.com/shortformikael/AlertHub/src/services/logger"
	"github.com/shortformikael/AlertHub/src/services/server"
)

const logComponent = "HANDLER"

func Start() {
	logger.Log(logComponent, logger.INFO, "Handler Started")
	defer logger.Log(logComponent, logger.INFO, "Handler ended")
	server.Init()
}
