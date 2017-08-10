package main

import (
	"fmt"
	"os"

	log "github.com/cihub/seelog"
	"github.com/grindlemire/seezlog"
)

func main() {
	err := seezlog.SetupLogger(seezlog.Critical, seezlog.Error, "./example.log")
	if err != nil {
		fmt.Printf("Error setting up logger: %v", err)
		os.Exit(1)
	}

	log.Info("Woo info logs go to stdout and the log file")
	log.Warn("Woo warn logs go to stdout and the log file")
	log.Error("Woo error logs go to stdout and the log file")
	log.Critical("Critical logs only go to the file")
	log.Debug("Debug logs are not logged")
	log.Trace("Trace logs are not logged either")
}
