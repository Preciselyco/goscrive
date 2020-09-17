package scrive

import (
	"fmt"
	"bytes"
	"encoding/json"
)

func getJsonDecoder(body []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	return dec
}

func parseJson(body []byte, out interface{}) error {
	return getJsonDecoder(body).Decode(out)
}

func logInvalidJson(body []byte) {
	if len(body) > 4096 {
		body = body[:4096]
	}
	fmt.Println(string(body))
}
