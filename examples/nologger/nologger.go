package main

import (
	"fmt"
	"os"

	log "github.com/cihub/seelog"
	seezlog "github.com/grindlemire/SeeZLog"
)

func main() {

	logger, err := seezlog.SetupLogger(seezlog.Trace, seezlog.NoLog, "./example.log")
	if err != nil {
		fmt.Printf("Error setting up logger: %v", err)
		os.Exit(1)
	}
	log.ReplaceLogger(logger)
	defer log.Flush()

	log.Info("All info logs should go to the file only")
	log.Warn("All warn logs should go to the file only")
	log.Error("All error logs should go to the file only")
	log.Critical("All critical logs should go to the file only")
	log.Debug("All debug logs should go to the file only")
	log.Trace("All trace logs should go to the file only")
}
