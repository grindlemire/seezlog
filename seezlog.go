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

type hasConsole int

const (
	withConsole = iota
	withoutConsole
	onlyConsole
)

type filterCtor func(path string) filter
type consoleToFilter map[hasConsole]filterCtor

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
		withConsole:    newInfoConsoleFilter,
		withoutConsole: newInfoFilter,
		onlyConsole:    newInfoConsoleOnlyFilter,
	},
	Warn: consoleToFilter{
		withConsole:    newWarnConsoleFilter,
		withoutConsole: newWarnFilter,
		onlyConsole:    newWarnConsoleOnlyFilter,
	},
	Error: consoleToFilter{
		withConsole:    newErrorConsoleFilter,
		withoutConsole: newErrorFilter,
		onlyConsole:    newErrorConsoleOnlyFilter,
	},
	Critical: consoleToFilter{
		withConsole:    newCriticalConsoleFilter,
		withoutConsole: newCriticalFilter,
		onlyConsole:    newCriticalConsoleOnlyFilter,
	},
	Debug: consoleToFilter{
		withConsole:    newDebugConsoleFilter,
		withoutConsole: newDebugFilter,
		onlyConsole:    newDebugConsoleOnlyFilter,
	},
	Trace: consoleToFilter{
		withConsole:    newTraceConsoleFilter,
		withoutConsole: newTraceFilter,
		onlyConsole:    newTraceConsoleOnlyFilter,
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
		RollingFile: &rollingFile{
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
		RollingFile: &rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: &console{},
	}
}

func newInfoConsoleOnlyFilter(path string) (f filter) {
	return filter{
		Levels:   infoLevel,
		FormatID: infoFormatID,
		Console:  &console{},
	}
}

func newWarnFilter(path string) filter {
	return filter{
		Levels:   warnLevel,
		FormatID: warnFormatID,
		RollingFile: &rollingFile{
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
		RollingFile: &rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: &console{},
	}
}

func newWarnConsoleOnlyFilter(path string) (f filter) {
	return filter{
		Levels:   warnLevel,
		FormatID: warnFormatID,
		Console:  &console{},
	}
}

func newErrorFilter(path string) filter {
	return filter{
		Levels:   errorLevel,
		FormatID: errorFormatID,
		RollingFile: &rollingFile{
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
		RollingFile: &rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: &console{},
	}
}

func newErrorConsoleOnlyFilter(path string) (f filter) {
	return filter{
		Levels:   errorLevel,
		FormatID: errorFormatID,
		Console:  &console{},
	}
}

func newCriticalFilter(path string) filter {
	return filter{
		Levels:   criticalLevel,
		FormatID: criticalFormatID,
		RollingFile: &rollingFile{
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
		RollingFile: &rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: &console{},
	}
}

func newCriticalConsoleOnlyFilter(path string) (f filter) {
	return filter{
		Levels:   criticalLevel,
		FormatID: criticalFormatID,
		Console:  &console{},
	}
}

func newDebugFilter(path string) filter {
	return filter{
		Levels:   debugLevel,
		FormatID: debugFormatID,
		RollingFile: &rollingFile{
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
		RollingFile: &rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: &console{},
	}
}

func newDebugConsoleOnlyFilter(path string) (f filter) {
	return filter{
		Levels:   debugLevel,
		FormatID: debugFormatID,
		Console:  &console{},
	}
}

func newTraceFilter(path string) filter {
	return filter{
		Levels:   traceLevel,
		FormatID: traceFormatID,
		RollingFile: &rollingFile{
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
		RollingFile: &rollingFile{
			Type:     "size",
			FileName: path,
			MaxSize:  20000000,
			MaxRolls: 5,
		},
		Console: &console{},
	}
}

func newTraceConsoleOnlyFilter(path string) (f filter) {
	return filter{
		Levels:   traceLevel,
		FormatID: traceFormatID,
		Console:  &console{},
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

// SetupLogger will setup a seelog logger that only outputs to stdout
// to a specific log level. Call ReplaceLogger with this logger from seelog to use this logger.
func SetupConsoleLogger(outTo Level) (logger log.LoggerInterface, err error) {
	config, err := GenerateConfig(NoLog, outTo, "")
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

	if logTo == NoLog {
		filters = onlyConsoleConfig(outTo)
	} else if outTo == NoLog {
		filters = onlyLogConfig(logTo, path)
	} else {
		filters = bothConfig(logTo, outTo, path)
	}

	seeLog.Outputs.Filters = filters

	sConf, err := xml.Marshal(seeLog)
	if err != nil {
		return "", err
	}

	return string(sConf), err
}

func onlyConsoleConfig(outTo Level) (filters []filter) {
	for _, level := range orderedLogList {
		var filter filter
		if level <= outTo {
			filterCtor := levelToFilter[level][onlyConsole]
			filter = filterCtor("")
			filters = append(filters, filter)
		}
	}

	return filters
}

func onlyLogConfig(logTo Level, path string) (filters []filter) {
	for _, level := range orderedLogList {
		var filter filter
		if level <= logTo {
			filterCtor := levelToFilter[level][withoutConsole]
			filter = filterCtor(path)
			filters = append(filters, filter)
		}
	}

	return filters
}

func bothConfig(logTo Level, outTo Level, path string) (filters []filter) {
	for _, level := range orderedLogList {
		var filter filter
		if level <= outTo {
			filterCtor := levelToFilter[level][withConsole]
			filter = filterCtor(path)
			filters = append(filters, filter)
		} else if level <= logTo {
			filterCtor := levelToFilter[level][withoutConsole]
			filter = filterCtor(path)
			filters = append(filters, filter)
		}
	}

	return filters
}
