package scrive

import (
	"bytes"
	"encoding/json"
)

func getJsonDecoder(body []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	dec.UseNumber()
	return dec
}

func parseJson(body []byte, out interface{}) error {
	return getJsonDecoder(body).Decode(out)
}
