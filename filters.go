package logger

func NewSeverityFilter(minLevel Severity) Filter {
	return NewFilterFunc(func(msg Message) (Message, bool) {
		if msg.Severity <= minLevel {
			return msg, true
		}
		return msg, false
	})
}
