package main

import (
	"fmt"
	"os"

	log "github.com/cihub/seelog"
	"github.com/grindlemire/seezlog"
)

func main() {
	logger, err := seezlog.SetupConsoleLogger(seezlog.Error)
	if err != nil {
		fmt.Printf("Error setting up logger: %v", err)
		os.Exit(1)
	}
	log.ReplaceLogger(logger)
	defer log.Flush()

	log.Info("Woo info logs only go to stdout")
	log.Warn("Woo warn logs only go to stdout")
	log.Error("Woo error logs only go to stdout")
	log.Critical("Critical logs are not logged")
	log.Debug("Debug logs are not logged")
	log.Trace("Trace logs are not logged")
}
