package logger

import (
	"fmt"
	"time"
)

type loggerimpl struct {
	writer Writer
}

func NewLogger(writer Writer) Logger {
	return &loggerimpl{
		writer: writer,
	}
}

func (l *loggerimpl) Hijack(defaultSeverity Severity) {
	NewHijack(l, defaultSeverity).Do()
}

func (l *loggerimpl) Log(s Severity, m string) {
	msg := Message{
		Time:     time.Now(),
		Severity: s,
		Message:  m,
		Extra:    make(MessageExtra),
	}
	l.writer.Write(msg)
}

func (l *loggerimpl) Logf(s Severity, m string, additional ...interface{}) {
	l.Log(s, fmt.Sprintf(m, additional...))
}

func (l *loggerimpl) Trace(m string)                        { l.Log(TRACE, m) }
func (l *loggerimpl) Tracef(m string, a ...interface{})     { l.Logf(TRACE, m, a...) }
func (l *loggerimpl) Debug(m string)                        { l.Log(DEBUG, m) }
func (l *loggerimpl) Debugf(m string, a ...interface{})     { l.Logf(DEBUG, m, a...) }
func (l *loggerimpl) Info(m string)                         { l.Log(INFO, m) }
func (l *loggerimpl) Infof(m string, a ...interface{})      { l.Logf(INFO, m, a...) }
func (l *loggerimpl) Notice(m string)                       { l.Log(NOTICE, m) }
func (l *loggerimpl) Noticef(m string, a ...interface{})    { l.Logf(NOTICE, m, a...) }
func (l *loggerimpl) Warning(m string)                      { l.Log(WARNING, m) }
func (l *loggerimpl) Warningf(m string, a ...interface{})   { l.Logf(WARNING, m, a...) }
func (l *loggerimpl) Error(m string)                        { l.Log(ERROR, m) }
func (l *loggerimpl) Errorf(m string, a ...interface{})     { l.Logf(ERROR, m, a...) }
func (l *loggerimpl) Critical(m string)                     { l.Log(CRITICAL, m) }
func (l *loggerimpl) Criticalf(m string, a ...interface{})  { l.Logf(CRITICAL, m, a...) }
func (l *loggerimpl) Alert(m string)                        { l.Log(ALERT, m) }
func (l *loggerimpl) Alertf(m string, a ...interface{})     { l.Logf(ALERT, m, a...) }
func (l *loggerimpl) Emergency(m string)                    { l.Log(EMERGENCY, m) }
func (l *loggerimpl) Emergencyf(m string, a ...interface{}) { l.Logf(EMERGENCY, m, a...) }

func (l *loggerimpl) WithData(key string, value interface{}) Logger {
	writer := NewWriteFilter(func(msg Message) (Message, bool) {
		msg.Extra[key] = value
		return msg, true
	}, l.writer)
	return NewLogger(writer)
}

func (l *loggerimpl) WithPrefix(p string) Logger {
	writer := NewWriteFilter(func(msg Message) (Message, bool) {
		msg.Message = p + msg.Message
		return msg, true
	}, l.writer)
	return NewLogger(writer)
}
