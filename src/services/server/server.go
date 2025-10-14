package server

import (
	"github.com/shortformikael/AlertHub/src/config"
	"github.com/shortformikael/AlertHub/src/services/logger"
)

const logComponent = "SERVER"

var server struct {
	apiroute string
}

func Init() error {
	logger.Log(logComponent, logger.LogTypeInfo, "Server Initializing...")

	server.apiroute = config.Name + "/api"

	logger.Log(logComponent, logger.LogTypeInfo, "Server Initialized!")
	return nil
}
