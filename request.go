package scrive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
)

type RequestMethod = string

const (
	methodGET  RequestMethod = "GET"
	methodPOST RequestMethod = "POST"
)

const noExpect = -1

type request struct {
	method         RequestMethod
	path           string
	headers        *map[string]string
	body           []byte
	expectCode     int
	out            interface{}
	binaryResponse bool

	_mr      *multipart.Writer
	_bodyBuf bytes.Buffer
	_qp      *url.Values
	_ses     []*ScriveError
}

func (c *Client) newRequest(method RequestMethod) *request {
	return &request{
		method:     method,
		headers:    &map[string]string{},
		body:       make([]byte, 0),
		expectCode: noExpect,
		_ses:       make([]*ScriveError, 0),
	}
}

func (c *Client) newPostRequest() *request {
	return c.newRequest(methodPOST)
}

func (c *Client) newGetRequest() *request {
	return c.newRequest(methodGET)
}

func (r *request) Path(path string) *request {
	r.path = path
	return r
}

func (r *request) Header(key, val string) *request {
	(*r.headers)[key] = val
	return r
}

func (r *request) Body(body []byte) *request {
	r.body = body
	return r
}

func (r *request) Expect(code int, out interface{}) *request {
	r.expectCode = code
	r.out = out
	return r
}

func (r *request) ExpectBinary(code int) *request {
	r.expectCode = code
	r.binaryResponse = true
	return r
}

func (r *request) ReadCloser() io.ReadCloser {
	return ioutil.NopCloser(
		bytes.NewReader(r.body),
	)
}

func (r *request) finalize() *ScriveError {
	if r.method == methodPOST {
		if r._mr != nil {
			if err := r._mr.Close(); err != nil {
				return localError(err)
			}
			(*r.headers)["Content-Type"] = r._mr.FormDataContentType()
			r.body = r._bodyBuf.Bytes()
		}
	}
	return nil
}

func (r *request) _ensureMultipart() {
	if r._mr == nil {
		r._bodyBuf = bytes.Buffer{}
		r._mr = multipart.NewWriter(&r._bodyBuf)
	}
}

func (r *request) _writeMVal(field string, value string) *ScriveError {
	r._ensureMultipart()
	if err := r._mr.WriteField(field, value); err != nil {
		err := localError(err)
		r._ses = append(r._ses, err)
		return err
	}
	return nil
}

func (r *request) writeMFile(field string, fileName string, file []byte) *ScriveError {
	if len(file) == 0 {
		return nil
	}
	r._ensureMultipart()
	fw, err := r._mr.CreateFormFile(field, fileName)
	if err != nil {
		err := localError(err)
		r._ses = append(r._ses, err)
		return err
	}
	if _, err := io.Copy(fw, bytes.NewReader(file)); err != nil {
		err := localError(err)
		r._ses = append(r._ses, err)
		return err
	}
	return nil
}

func (r *request) writeMBool(field string, value *bool) *ScriveError {
	if value == nil {
		return nil
	}
	return r._writeMVal(field, boolToStr(*value))
}

func (r *request) writeMString(field string, value *string) *ScriveError {
	if value == nil {
		return nil
	}
	return r._writeMVal(field, *value)
}

func (r *request) writeMUInt(field string, value *uint64) *ScriveError {
	if value == nil {
		return nil
	}
	return r._writeMVal(field, fmt.Sprintf("%d", *value))
}

func (r *request) writeMStrdef(field string, value strDef) *ScriveError {
	if value == nil {
		return nil
	}
	return r._writeMVal(field, *(value.strp()))
}

func (r *request) writeMJSON(field string, obj interface{}) *ScriveError {
	b, err := json.Marshal(obj)
	if err != nil {
		return localError(err)
	}
	str := string(b)
	return r.writeMString(field, &str)
}

func (r *request) _ensureQuery() {
	if r._qp == nil {
		r._qp = &url.Values{}
	}
}

func (r *request) _addQuery(key string, value string) *ScriveError {
	r._ensureQuery()
	r._qp.Add(key, value)
	return nil
}

func (r *request) _setQuery(key string, value string) *ScriveError {
	r._ensureQuery()
	r._qp.Set(key, value)
	return nil
}

func (r *request) addQueryUint(key string, value *uint64) *ScriveError {
	if value == nil {
		return nil
	}
	return r._addQuery(key, fmt.Sprintf("%d", *value))
}

func (r *request) setQueryUint(key string, value *uint64) *ScriveError {
	if value == nil {
		return nil
	}
	return r._setQuery(key, fmt.Sprintf("%d", *value))
}

func (r *request) addQueryJSON(key string, value interface{}) *ScriveError {
	if value == nil {
		return nil
	}
	b, err := json.Marshal(value)
	if err != nil {
		return localError(err)
	}
	return r._addQuery(key, string(b))
}

func (r *request) setQueryJSON(key string, value interface{}) *ScriveError {
	if value == nil {
		return nil
	}
	b, err := json.Marshal(value)
	if err != nil {
		return localError(err)
	}
	return r._setQuery(key, string(b))
}

func (r *request) getQuery() string {
	if r._qp == nil {
		return ""
	}
	return r._qp.Encode()
}

func (r *request) anyError(errs ...*ScriveError) *ScriveError {
	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	return nil
}

func (r *request) anyErrors() *ScriveError {
	return r.anyError(r._ses...)
}
