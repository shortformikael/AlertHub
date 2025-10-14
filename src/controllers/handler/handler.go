package handler

import (
	"github.com/shortformikael/AlertHub/src/services/logger"
)

const logComponent = "HANDLER"

func Start() {
	logger.Log(logComponent, logger.LogTypeInfo, "Handler Started")

	defer logger.Log(logComponent, logger.LogTypeInfo, "Handler ended")
}
