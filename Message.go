package logger

import (
	"time"
)

type Message struct {
	Time     time.Time    `json:"time"`
	Severity Severity     `json:"severity"`
	Message  string       `json:"message,omitempty"`
	Extra    MessageExtra `json:"extra,omitempty"`
}
