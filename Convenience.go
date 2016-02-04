package logger

import (
	"io"
	"os"
)

func NewStringLogger(o io.Writer) *Logger {
	return NewLogger(
		NewFileWriter(
			o,
			NewStringEncoder(),
			Separators.NewLine,
		),
	)
}

func NewStdoutLogger() *Logger {
	return NewStringLogger(os.Stdout)
}

func NewStderrLogger() *Logger {
	return NewStringLogger(os.Stderr)
}
