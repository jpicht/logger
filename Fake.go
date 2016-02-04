package logger

func Fake() *Logger {
	return NewLogger(&nullwriter{})
}

type nullwriter struct {
}

func (n *nullwriter) Write(Message) {
}
