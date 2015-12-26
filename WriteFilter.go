package logger

type WriteFilterFn func(Message) (Message, bool)

type WriteFilter struct {
	fn     WriteFilterFn
	writer Writer
}

func NewWriteFilter(fn WriteFilterFn, writer Writer) *WriteFilter {
	return &WriteFilter{
		fn:     fn,
		writer: writer,
	}
}

func (w *WriteFilter) Write(msg Message) {
	if m, ok := w.fn(msg); ok {
		w.writer.Write(m)
	}
}
