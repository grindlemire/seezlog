package seezlog

type seelog struct {
	Type          string   `xml:"type,attr"`
	AsyncInterval int      `xml:"asyncinterval,attr"`
	Outputs       outputs  `xml:"outputs"`
	Formats       []format `xml:"formats>format"`
}

type outputs struct {
	Filters []filter `xml:"filter"`
}

type format struct {
	ID     string `xml:"id,attr"`
	Format string `xml:"format,attr"`
}

type filter struct {
	Levels      string       `xml:"levels,attr"`
	FormatID    string       `xml:"formatid,attr"`
	Console     *console     `xml:"console,omitempty"`
	RollingFile *rollingFile `xml:"rollingfile,omitempty"`
}

type console struct{}

type rollingFile struct {
	Type     string `xml:"type,attr,omitempty"`
	FileName string `xml:"filename,attr,omitempty"`
	MaxSize  int    `xml:"maxsize,attr,omitempty"`
	MaxRolls int    `xml:"maxrolls,attr,omitempty"`
}

// Format strings that define the way the log is printed to the log
const (
	infoFmtStr     = "%EscM(32)[%Level]%EscM(0) [%Date %Time] [%File] %Msg%n"
	warnFmtStr     = "%EscM(33)[%LEVEL]%EscM(0) [%Date %Time] [%FuncShort @ %File.%Line] %Msg%n"
	errorFmtStr    = "%EscM(31)[%LEVEL]%EscM(0) [%Date %Time] [%FuncShort @ %File.%Line] %Msg%n"
	criticalFmtStr = "%EscM(31)[%LEVEL]%EscM(0) [%Date %Time] [%FuncShort @ %File.%Line] %Msg%n"
	debugFmtStr    = "%EscM(34)[%LEVEL]%EscM(0) [%Date %Time] [%FuncShort @ %File:%Line] %Msg%n"
	traceFmtStr    = "%EscM(36)[%LEVEL]%EscM(0) [%Date %Time] [%FuncShort @ %File:%Line] %Msg%n"
)

// Levels indicating which filter applies to which log level
const (
	infoLevel     = "info"
	warnLevel     = "warn"
	errorLevel    = "error"
	criticalLevel = "critical"
	debugLevel    = "debug"
	traceLevel    = "trace"
)

// format ids for every level of logging
const (
	infoFormatID     = "fmtinfo"
	warnFormatID     = "fmtwarn"
	errorFormatID    = "fmterror"
	criticalFormatID = "fmtcritical"
	debugFormatID    = "fmtdebug"
	traceFormatID    = "fmttrace"
)

// Format objects that define a seelog format
var (
	infoFmt = format{
		ID:     infoFormatID,
		Format: infoFmtStr,
	}

	warnFmt = format{
		ID:     warnFormatID,
		Format: warnFmtStr,
	}

	errorFmt = format{
		ID:     errorFormatID,
		Format: errorFmtStr,
	}

	criticalFmt = format{
		ID:     criticalFormatID,
		Format: criticalFmtStr,
	}

	debugFmt = format{
		ID:     debugFormatID,
		Format: debugFmtStr,
	}

	traceFmt = format{
		ID:     traceFormatID,
		Format: traceFmtStr,
	}

	allFmts = []format{
		infoFmt,
		warnFmt,
		errorFmt,
		criticalFmt,
		debugFmt,
		traceFmt,
	}
)
