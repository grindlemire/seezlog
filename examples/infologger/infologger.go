package main

import (
	"fmt"
	"os"

	log "github.com/cihub/seelog"
	"github.com/grindlemire/seezlog"
)

func main() {
	logger, err := seezlog.SetupLogger(seezlog.Info, seezlog.Critical, "./example.log")
	if err != nil {
		fmt.Printf("Error setting up logger: %v", err)
		os.Exit(1)
	}
	log.ReplaceLogger(logger)
	defer log.Flush()

	log.Critical("Woo critical logs go to stdout and the log file")
	log.Error("Error logs only go to the file")
	log.Warn("Warn logs only go to the file")
	log.Info("Info logs only go to the file")
	log.Debug("Debug logs are not logged")
	log.Trace("Trace logs are not logged either")
}
