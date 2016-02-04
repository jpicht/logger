package logger

import (
	ctx "golang.org/x/net/context"

	"errors"
)

type loggerCtxKey string

const loggerKey = loggerCtxKey("logger")

var NoLoggerInContext = errors.New("Context does not provide a logger instance")

func (l *Logger) Context(p ctx.Context) ctx.Context {
	return ctx.WithValue(p, loggerKey, l)
}

func FromContext(c ctx.Context) (*Logger, error) {
	l_i := c.Value(loggerKey)
	if l, ok := l_i.(*Logger); ok {
		return l, nil
	}
	return nil, NoLoggerInContext
}

func MayFromContext(c ctx.Context) *Logger {
	l, err := FromContext(c)
	if nil != err {
		return Fake()
	}
	return l
}

func MustFromContext(c ctx.Context) *Logger {
	l, err := FromContext(c)
	if nil != err {
		panic(err)
	}
	return l
}
