package logger

import (
	"io"
)

type FileWriter struct {
	writer  io.Writer
	encoder Encoder
}

func NewFileWriter(out io.Writer) *FileWriter {
	return &FileWriter{
		writer:  out,
		encoder: NewStringEncoder(),
	}
}

func (f *FileWriter) Write(m Message) {
	f.writer.Write(f.encoder.Encode(m))
	f.writer.Write([]byte("\n"))
}
