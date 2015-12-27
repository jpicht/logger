package logger

import (
	//	"io"
	"testing"
)

type MockWriter struct {
	buff []byte
}

func NewMockWriter() *MockWriter {
	return &MockWriter{
		buff: make([]byte, 0),
	}
}

func (m *MockWriter) Write(in []byte) (n int, err error) {
	m.buff = append(m.buff, in...)
	return len(in), nil
}

func (m *MockWriter) Flush() {
	m.buff = make([]byte, 0)
}

func (m *MockWriter) Get() []byte {
	return m.buff
}

func TestFileWriter(t *testing.T) {
	doTest := func(str string) {
		mw := NewMockWriter()
		se := NewStringEncoder()
		se.Format = "%M"
		fw := NewFileWriter(mw, se, Separators.Empty)
		fw.Write(Message{Message: str})
		if string(mw.Get()) != str {
			t.Fatalf("Wrong output: '%s' expected '%s'", mw.Get(), str)
		}
	}
	doTest("test")
	doTest("")
	doTest("\n")
}
