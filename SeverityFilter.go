package logger

func NewSeverityFilter(min Severity, writer Writer) Writer {
	return NewWriteFilter(func(msg Message) (Message, bool) {
		return msg, msg.Severity <= min
	}, writer)
}
