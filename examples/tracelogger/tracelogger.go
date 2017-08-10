package main

import (
	"fmt"
	"os"

	log "github.com/cihub/seelog"
	seezlog "github.com/grindlemire/SeeZLog"
)

func main() {

	err := seezlog.SetupLogger(seezlog.Trace, seezlog.Trace, "./example.log")
	if err != nil {
		fmt.Printf("Error setting up logger: %v", err)
		os.Exit(1)
	}
	defer log.Flush()

	log.Info("Woo info logs go to stdout and the log file")
	log.Warn("Woo warn logs go to stdout and the log file")
	log.Error("Woo error logs go to stdout and the log file")
	log.Critical("Woo critical logs go to stdout and the log file")
	log.Debug("Woo debug logs go to stdout and the log file")
	log.Trace("Woo trace logs go to stdout and the log file")
}
