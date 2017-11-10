# seezlog
seezlog is a convenience wrapper for [seelog](https://github.com/cihub/seelog) so you can setup formatted logging in one line and it will "just work".

[Godoc](http://godoc.org/github.com/Grindlemire/seezlog)


# Why?
Because I found myself wanting to use seelog as a logger for command line tools but found it tedious to set up the xml formatting in every project. This will wrap all that logic and provide you a one line function that initializes a seelog logger capable of printing both to stdout and a log file.


# Usage
For simplicity, there are only three public functions:

### `SetupLogger`
```Go
SetupLogger(logTo Level, outTo Level, path string) (logger log.LoggerInterface, err error)
```
Creates and initializes up a seelog logger that will print to stdout up to `outTo` log level and will print to a log file up to `logTo` log level. Just plug this into seelog's `ReplaceLogger` function to use it. See the example below or take a look in the [examples](https://github.com/Grindlemire/seezlog/tree/master/examples) directory

---
### `SetupConsoleLogger`
```Go
SetupConsoleLogger(outTo Level) (logger log.LoggerInterface, err error)
```
Creates and initializes up a seelog logger that will only print to stdout up to `outTo` log level. Just plug this into seelog's `ReplaceLogger` function to use it. See the example below or take a look in the [examples](https://github.com/Grindlemire/seezlog/tree/master/examples) directory

---
### `GenerateConfig`
```Go
GenerateConfig(logTo Level, outTo Level, path string) (config string, err error)
```
Generates the configuration to create a seelog logger that will print to stdout up to `outTo` log level and will print to a log file up to `logTo` log level. This can be useful if you are trying to generate a constant config file for seelog and just want this library to generate a no hastle config for you.

### Log  Levels
The accepted log levels of type `Level` are
`Critical`, `Error`, `Warn`, `Info`, `Debug`, `Trace`, and `NoLog`

The level ordering is as follows: `Critical` > `Error` > `Warn` > `Info` > `Debug` > `Trace`

---

### Example
```Go

import(
    seezlog "github.com/grindlemire/seezlog"
    log "github.com/cihub/seelog"
)

logger, err := seezlog.SetupLogger(seezlog.Info, seezlog.Warn, "./example.log")
if err != nil {
    fmt.Printf("Error setting up logger: %v", err)
    os.Exit(1)
}
log.ReplaceLogger(logger)
defer log.Flush()

log.Critical("This critical will to stdout and the log file")
log.Warn("This warn will print to stdout and the log file")
log.Error("This error will print only to the log file")
log.Info("This info will print only to the log file")
log.Debug("This debug will not print to the log or stdout")
log.Trace("This trace will not print to the log or stdout")
```


# Example Log Messages

<span style="color:green">[Info]</span> [2006-01-02 15:04:05] [file.go] Info message

<span style="color:yellow">[WARN]</span> [2006-01-02 15:04:05] [func @ file.go.Line] Warn message

<span style="color:red">[ERROR]</span> [2006-01-02 15:04:05] [func @ file.go.Line] Error message

<span style="color:darkred">[CRITICAL]</span> [2006-01-02 15:04:05] [func @ file.go.Line] Critical message

<span style="color:blue">[DEBUG]</span> [2006-01-02 15:04:05] [func @ file.go:Line] Debug message

<span style="color:cyan">[TRACE]</span> [2006-01-02 15:04:05] [func @ file.go:Line] Trace message
