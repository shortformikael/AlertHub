package server

import (
	"github.com/shortformikael/AlertHub/src/config"
	"github.com/shortformikael/AlertHub/src/services/logger"
)

const logComponent = "SERVER"

var server struct {
	apiroute string
}

func Init() {
	logger.Log(logComponent, logger.INFO, "Server Initializing...")
	defer logger.Log(logComponent, logger.INFO, "Server Initialized!")
	server.apiroute = config.Name + "/api"
}
