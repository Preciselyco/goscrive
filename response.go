package scrive

import (
	"net/http"
	"strings"
)

type response struct {
	code    int
	headers *map[string]string
	body    []byte
}

func newResponse(code int, header http.Header, body []byte) *response {
	resp := &response{
		code:    code,
		headers: &map[string]string{},
	}

	for k, v := range header {
		(*resp.headers)[strings.ToLower(k)] = strings.Join(v, ",")
	}
	bLen := len(body)
	if bLen > 0 {
		resp.body = make([]byte, bLen)
		copy(resp.body, body)
	}
	return resp
}

func (r *response) getHeader(k string) string {
	return (*r.headers)[k]
}
