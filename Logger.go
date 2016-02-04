package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	writer Writer
}

func NewLogger(writer Writer) *Logger {
	return &Logger{
		writer: writer,
	}
}

func (l *Logger) Hijack(defaultSeverity Severity) {
	NewHijack(l, defaultSeverity).Do()
}

func (l *Logger) Log(s Severity, m string) {
	msg := Message{
		Time:     time.Now(),
		Severity: s,
		Message:  m,
		Extra:    make(MessageExtra),
	}
	l.writer.Write(msg)
}

func (l *Logger) Logf(s Severity, m string, additional ...interface{}) {
	l.Log(s, fmt.Sprintf(m, additional...))
}

func (l *Logger) Trace(m string)                        { l.Log(TRACE, m) }
func (l *Logger) Tracef(m string, a ...interface{})     { l.Logf(TRACE, m, a...) }
func (l *Logger) Debug(m string)                        { l.Log(DEBUG, m) }
func (l *Logger) Debugf(m string, a ...interface{})     { l.Logf(DEBUG, m, a...) }
func (l *Logger) Info(m string)                         { l.Log(INFO, m) }
func (l *Logger) Infof(m string, a ...interface{})      { l.Logf(INFO, m, a...) }
func (l *Logger) Notice(m string)                       { l.Log(NOTICE, m) }
func (l *Logger) Noticef(m string, a ...interface{})    { l.Logf(NOTICE, m, a...) }
func (l *Logger) Warning(m string)                      { l.Log(WARNING, m) }
func (l *Logger) Warningf(m string, a ...interface{})   { l.Logf(WARNING, m, a...) }
func (l *Logger) Error(m string)                        { l.Log(ERROR, m) }
func (l *Logger) Errorf(m string, a ...interface{})     { l.Logf(ERROR, m, a...) }
func (l *Logger) Critical(m string)                     { l.Log(CRITICAL, m) }
func (l *Logger) Criticalf(m string, a ...interface{})  { l.Logf(CRITICAL, m, a...) }
func (l *Logger) Alert(m string)                        { l.Log(ALERT, m) }
func (l *Logger) Alertf(m string, a ...interface{})     { l.Logf(ALERT, m, a...) }
func (l *Logger) Emergency(m string)                    { l.Log(EMERGENCY, m) }
func (l *Logger) Emergencyf(m string, a ...interface{}) { l.Logf(EMERGENCY, m, a...) }

func (l *Logger) WithData(key string, value interface{}) *Logger {
	writer := NewWriteFilter(func(msg Message) (Message, bool) {
		msg.Extra[key] = value
		return msg, true
	}, l.writer)
	return NewLogger(writer)
}

func (l *Logger) WithPrefix(p string) *Logger {
	writer := NewWriteFilter(func(msg Message) (Message, bool) {
		msg.Message = p + msg.Message
		return msg, true
	}, l.writer)
	return NewLogger(writer)
}
