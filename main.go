package main

import (
	"github.com/DerryRenaldy/learnFiber/server"
	"github.com/DerryRenaldy/logger/constant"
	"github.com/DerryRenaldy/logger/logger"
)

func main() {
	logger := logger.New("fiber test", constant.EnvDev, constant.DebugLevel)
	sv := server.New(logger)

	sv.Start()
}
