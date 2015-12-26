package logger

type MultiWriter struct {
	writers []Writer
}

func NewMultiWriter(w ...Writer) *MultiWriter {
	writers := make([]Writer, 0)
	writers = append(writers, w...)
	return &MultiWriter{
		writers: writers,
	}
}

func (m *MultiWriter) Add(w ...Writer) {
	m.writers = append(m.writers, w...)
}

func (m *MultiWriter) Write(msg Message) {
	for _, w := range m.writers {
		w.Write(msg)
	}
}
