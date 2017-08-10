package seezlog

import (
	"encoding/xml"
	"fmt"

	log "github.com/cihub/seelog"
)

// Level is the enum type representing the log level
type Level int

// Log levels accepted by GenerateConfig and SetupLogger.
const (
	NoLog = iota
	Info
	Warn
	Error
	Critical
	Debug
	Trace
)

type withConsole int

const (
	with = iota
	without
)

type filterCtor func(path string) filter
type consoleToFilter map[withConsole]filterCtor

var orderedLogList = []Level{
	Info,
	Warn,
	Error,
	Critical,
	Debug,
	Trace,
}

var levelToFilter = map[Level]consoleToFilter{
	Info: consoleToFilter{
		with:    newInfoConsoleFilter,
		without: newInfoFilter,
	},
	Warn: consoleToFilter{
		with:    newWarnConsoleFilter,
		without: newWarnFilter,
	},
	Error: consoleToFilter{
		with:    newErrorConsoleFilter,
		without: newErrorFilter,
	},
	Critical: consoleToFilter{
		with:    newCriticalConsoleFilter,
		without: newCriticalFilter,
	},
	Debug: consoleToFilter{
		with:    newDebugConsoleFilter,
		without: newDebugFilter,
	},
	Trace: consoleToFilter{
		with:    newTraceConsoleFilter,
		without: newTraceFilter,
	},
}

// Seelog common object
func newCommonSeelog() seelog {
	return seelog{
		Type:          "asynctimer",
		AsyncInterval: 1000000,
		Formats:       allFmts,
	}
}

func newInfoFilter(path string) filter {
	return filter{
		Levels:   infoLevel,
		FormatID: infoFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
	}
}

func newInfoConsoleFilter(path string) (f filter) {
	return filter{
		Levels:   infoLevel,
		FormatID: infoFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: true,
	}
}

func newWarnFilter(path string) filter {
	return filter{
		Levels:   warnLevel,
		FormatID: warnFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
	}
}

func newWarnConsoleFilter(path string) (f filter) {
	return filter{
		Levels:   warnLevel,
		FormatID: warnFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: true,
	}
}

func newErrorFilter(path string) filter {
	return filter{
		Levels:   errorLevel,
		FormatID: errorFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
	}
}

func newErrorConsoleFilter(path string) (f filter) {
	return filter{
		Levels:   errorLevel,
		FormatID: errorFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: true,
	}
}

func newCriticalFilter(path string) filter {
	return filter{
		Levels:   criticalLevel,
		FormatID: criticalFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
	}
}

func newCriticalConsoleFilter(path string) (f filter) {
	return filter{
		Levels:   criticalLevel,
		FormatID: criticalFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: true,
	}
}

func newDebugFilter(path string) filter {
	return filter{
		Levels:   debugLevel,
		FormatID: debugFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
	}
}

func newDebugConsoleFilter(path string) (f filter) {
	return filter{
		Levels:   debugLevel,
		FormatID: debugFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: true,
	}
}

func newTraceFilter(path string) filter {
	return filter{
		Levels:   traceLevel,
		FormatID: traceFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
	}
}

func newTraceConsoleFilter(path string) (f filter) {
	return filter{
		Levels:   traceLevel,
		FormatID: traceFormatID,
		RollingFile: rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: true,
	}
}

// SetupLogger will setup a seelog logger to a specific log level and will output to stdout
// to a specific log level. Call ReplaceLogger with this logger from seelog to use this logger.
func SetupLogger(logTo Level, outTo Level, path string) (logger log.LoggerInterface, err error) {
	config, err := GenerateConfig(logTo, outTo, path)
	if err != nil {
		return nil, err
	}

	logger, err = log.LoggerFromConfigAsBytes([]byte(config))
	if err != nil {
		return nil, fmt.Errorf("error configuring logger from inline xml: %v", err)
	}
	return logger, nil
}

// GenerateConfig will generate the seelog xml file and spit it out to stdout.
func GenerateConfig(logTo Level, outTo Level, path string) (config string, err error) {
	seeLog := newCommonSeelog()
	filters := []filter{}
	for _, level := range orderedLogList {
		var filter filter
		if outTo != NoLog && level <= outTo {
			filterCtor := levelToFilter[level][with]
			filter = filterCtor(path)
			filters = append(filters, filter)
		} else if logTo != NoLog && level <= logTo {
			filterCtor := levelToFilter[level][without]
			filter = filterCtor(path)
			filters = append(filters, filter)
		}
	}
	seeLog.Outputs.Filters = filters

	sConf, err := xml.Marshal(seeLog)
	if err != nil {
		return "", err
	}

	return string(sConf), err
}
