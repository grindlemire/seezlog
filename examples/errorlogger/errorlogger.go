package main

import (
	"fmt"
	"os"

	log "github.com/cihub/seelog"
	"github.com/grindlemire/seezlog"
)

func main() {
	logger, err := seezlog.SetupLogger(seezlog.Error, seezlog.Critical, "./example.log")
	if err != nil {
		fmt.Printf("Error setting up logger: %v", err)
		os.Exit(1)
	}
	log.ReplaceLogger(logger)
	defer log.Flush()

	log.Critical("Woo critical logs go to stdout and the log file")
	log.Error("Error logs go to just the log")
	log.Warn("Warn logs go to just the log")
	log.Info("Info logs go to just the log")
	log.Debug("Debug logs are not logged")
	log.Trace("Trace logs are not logged either")
}
