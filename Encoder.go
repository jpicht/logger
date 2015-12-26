package logger

type Encoder interface {
	Encode(Message) []byte
}
