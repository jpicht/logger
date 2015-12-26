package logger

type FilterFunc func(Message) (Message, bool)

type filterFuncWrapper struct {
	fn func(Message) (Message, bool)
}

func NewFilterFunc(ff FilterFunc) Filter {
	return &filterFuncWrapper{
		fn: ff,
	}
}

func (ff *filterFuncWrapper) Filter(msg Message) (Message, bool) {
	nmsg, ok := ff.fn(msg)
	return nmsg, ok
}
