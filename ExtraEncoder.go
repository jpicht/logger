package logger

type ExtraEncoder interface {
	Encode(MessageExtra) []byte
}
