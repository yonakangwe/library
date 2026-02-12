package util

import (
	"bytes"

	"encoding/json"
)

func Decode(in, out interface{}) interface{} {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
	return out
}
