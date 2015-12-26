package logger

import "fmt"

type Severity int

const (
	EMERGENCY = Severity(0)
	ALERT     = Severity(1)
	CRITICAL  = Severity(2)
	ERROR     = Severity(3)
	WARNING   = Severity(4)
	NOTICE    = Severity(5)
	INFO      = Severity(6)
	DEBUG     = Severity(7)
	TRACE     = Severity(8)
)

func (s Severity) String() string {
	switch s {
	case EMERGENCY:
		return "EMERGENCY"
	case ALERT:
		return "ALERT"
	case CRITICAL:
		return "CRITICAL"
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case NOTICE:
		return "NOTICE"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	}
	return fmt.Sprintf("INVALID(%d)", int(s))
}
