# seezlog
seezlog is a convenience wrapper for [seelog](https://github.com/cihub/seelog) so you can setup formatted logging in one line and it will "just work".

Look at the [Godoc](http://godoc.org/github.com/grindlemire/seezlog) for full documentation.


# Why?
Because I found myself wanting to use seelog as a logger for command line tools but found it tedious to set up the xml formatting in every project. This will wrap all that logic and provide you a one line function that initializes a seelog logger capable of printing both to stdout and a log file.


# Usage
There are only two public functions to keep this simple to use:

### `SetupLogger`
```Go
SetupLogger(logTo Level, outTo Level, path string) (logger log.LoggerInterface, err error)
```
Creates and initializes up a seelog logger that will print to stdout up to `outTo` level and will print to a log file up to `logTo` level. Just plug this into seelog's `ReplaceLogger` function to use it. See the example below or take a look in the [examples](https://github.com/Grindlemire/seezlog/tree/master/examples) directory

---
### `GenerateConfig`
```Go
GenerateConfig(logTo Level, outTo Level, path string) (config string, err error)
```
Generates the configuration to create a seelog logger that will print to stdout up to `outTo` level and will print to a log file up to `logTo` level. This can be useful if you are trying to generate a const config file for seelog and just want this library to generate a no hastle config for you.

### Log  Levels
The accepted log levels of type `Level` are
`Info`, `Warn`, `Error`, `Critical`, `Debug`, `Trace`, and `NoLog`

---

### Example
```Go

import(
    seezlog "github.com/grindlemire/seezlog"
    log "github.com/cihub/seelog"
)

logger, err := seezlog.SetupLogger(seezlog.Critical, seezlog.Warn, "./example.log")
if err != nil {
    fmt.Printf("Error setting up logger: %v", err)
    os.Exit(1)
}
log.ReplaceLogger(logger)
defer log.Flush()

log.Info("This info will print to stdout and the log file")
log.Warn("This warn will print to stdout and the log file")
log.Error("This error will print only to the log file")
log.Critical("This critical will print only to the log file")
log.Debug("This debug will print only to the log file")
log.Trace("This trace will print only to the log file")
```