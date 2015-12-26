package logger

type Filter interface {
	Filter(Message) (Message, bool)
}
