package logger

import (
	"time"
)

type Message struct {
	Time     time.Time
	Severity Severity
	Message  string
	Extra    MessageExtra
}
