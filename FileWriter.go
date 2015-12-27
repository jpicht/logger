package logger

import (
	"io"
)

type FileWriter struct {
	writer    io.Writer
	encoder   Encoder
	separator []byte
}

type seperators struct {
	Empty   []byte
	NewLine []byte
}

var (
	Seperators = seperators{
		Empty:   make([]byte, 0),
		NewLine: []byte("\n"),
	}
)

func NewFileWriter(out io.Writer, e Encoder, separator []byte) *FileWriter {
	if nil == e {
		e = NewStringEncoder()
	}
	return &FileWriter{
		writer:    out,
		encoder:   e,
		separator: separator,
	}
}

func (f *FileWriter) Write(m Message) {
	f.writer.Write(f.encoder.Encode(m))
	f.writer.Write(f.separator)
}
