package loglevel

// LogLevel represents the severity of the log message.
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// logLevelToString converts a LogLevel to its string representation.
func LogLevelToString(level LogLevel) string {
	if level == DEBUG {
		return "DEBUG"
	}
	if level == INFO {
		return "INFO"
	}
	if level == WARN {
		return "WARN"
	}
	if level == ERROR {
		return "ERROR"
	}
	if level == FATAL {
		return "FATAL"
	}
	return "UNKNOWN"
}
