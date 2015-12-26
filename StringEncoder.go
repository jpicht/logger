package logger

import (
	"strings"
)

type StringEncoder struct {
	Format       string
	DateFormat   string
	ExtraEncoder ExtraEncoder
}

func NewStringEncoder() StringEncoder {
	return StringEncoder{
		Format:       "%d [%S]: %M%E",
		DateFormat:   "2006-01-02 15:04:05.000000",
		ExtraEncoder: NewStringExtraEncoder(),
	}
}

func (s StringEncoder) Encode(m Message) []byte {
	out := s.Format
	date := m.Time.Format(s.DateFormat)
	out = strings.Replace(out, "%d", date, 1)
	out = strings.Replace(out, "%S", m.Severity.String(), 1)
	out = strings.Replace(out, "%M", m.Message, 1)
	out = strings.Replace(out, "%E", string(s.ExtraEncoder.Encode(m.Extra)), 1)
	return []byte(out)
}
