package logger

import (
	ctx "golang.org/x/net/context"
)

type Logger interface {
	Context(p ctx.Context) ctx.Context
	Hijack(defaultSeverity Severity)

	Log(s Severity, m string)
	Logf(s Severity, m string, additional ...interface{})

	Trace(m string)
	Tracef(m string, a ...interface{})
	Debug(m string)
	Debugf(m string, a ...interface{})
	Info(m string)
	Infof(m string, a ...interface{})
	Notice(m string)
	Noticef(m string, a ...interface{})
	Warning(m string)
	Warningf(m string, a ...interface{})
	Error(m string)
	Errorf(m string, a ...interface{})
	Critical(m string)
	Criticalf(m string, a ...interface{})
	Alert(m string)
	Alertf(m string, a ...interface{})
	Emergency(m string)
	Emergencyf(m string, a ...interface{})

	WithData(key string, value interface{}) Logger
	WithPrefix(p string) Logger
}
