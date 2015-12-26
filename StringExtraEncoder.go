package logger

import (
	"fmt"
	"strings"
)

type StringExtraEncoder struct {
	OuterFormat string
	ItemFormat  string
	Glue        string
}

func NewStringExtraEncoder() *StringExtraEncoder {
	return &StringExtraEncoder{
		OuterFormat: " [%s]",
		ItemFormat:  "%s: %v",
		Glue:        ", ",
	}
}

func (s *StringExtraEncoder) Encode(extra MessageExtra) []byte {
	if len(extra) < 1 {
		return []byte("")
	}
	output := make([]string, len(extra))
	pos := 0
	for key, value := range extra {
		output[pos] = fmt.Sprintf(s.ItemFormat, key, value)
		pos += 1
	}
	output_str := fmt.Sprintf(s.OuterFormat, strings.Join(output, s.Glue))
	return []byte(output_str)
}
