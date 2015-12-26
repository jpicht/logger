package logger

import (
	"time"
)

type Logger struct {
	writer  Writer
	filters []Filter
}

func NewLogger(writer Writer) *Logger {
	return &Logger{
		writer:  writer,
		filters: make([]Filter, 0),
	}
}

func (l *Logger) AddFilter(f Filter) {
	l.filters = append(l.filters, f)
}

func (l *Logger) Log(m string, s Severity) {
	msg := Message{
		Time:     time.Now(),
		Severity: s,
		Message:  m,
	}
	for _, f := range l.filters {
		var ok bool
		msg, ok = f.Filter(msg)
		if !ok {
			return
		}
	}
	l.writer.Write(msg)
}

func (l *Logger) Trace(m string)     { l.Log(m, TRACE) }
func (l *Logger) Debug(m string)     { l.Log(m, DEBUG) }
func (l *Logger) Info(m string)      { l.Log(m, INFO) }
func (l *Logger) Notice(m string)    { l.Log(m, NOTICE) }
func (l *Logger) Warning(m string)   { l.Log(m, WARNING) }
func (l *Logger) Error(m string)     { l.Log(m, ERROR) }
func (l *Logger) Critical(m string)  { l.Log(m, CRITICAL) }
func (l *Logger) Alert(m string)     { l.Log(m, ALERT) }
func (l *Logger) Emergency(m string) { l.Log(m, EMERGENCY) }
