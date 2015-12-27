package logger

import (
	"encoding/json"
)

type JsonEncoder struct {
}

func NewJsonEncoder() JsonEncoder {
	return JsonEncoder{}
}

func (j JsonEncoder) Encode(m Message) []byte {
	j_text, err := json.Marshal(m)
	if nil != err {
		return []byte(err.Error())
	}
	return j_text
}
