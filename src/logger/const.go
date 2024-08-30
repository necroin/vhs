package logger

const (
	NoneLevel = iota
	ErrorLevel
	InfoLevel
	VerboseLevel
	DebugLevel
	DetailedLevel
)

var (
	levelMap = map[string]int{
		"error":    ErrorLevel,
		"info":     InfoLevel,
		"verbose":  VerboseLevel,
		"debug":    DebugLevel,
		"detailed": DetailedLevel,
	}
)
