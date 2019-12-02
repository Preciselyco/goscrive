package scrive

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

const (
	defaultAPIRoot      = "api-testbed.scrive.com"
	headerAuthorization = "Authorization"
)

type Config struct {
	APIRoot *string
	PAC     *PAC
}

type Client struct {
	config Config
	Debug  bool
}

func NewClient(c Config) (*Client, error) {
	client := &Client{
		config: Config{},
	}
	if c.PAC == nil {
		return nil, fmt.Errorf("Only authentication with PAC is currently supported")
	}
	client.config.PAC = c.PAC
	if c.APIRoot == nil {
		apiRoot := defaultAPIRoot
		client.config.APIRoot = &apiRoot
	} else {
		client.config.APIRoot = c.APIRoot
	}
	return client, nil
}

func (c *Client) httpClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}

func (c *Client) composeURL(path string) string {
	return fmt.Sprintf("https://%s/api/v2/%s", *c.config.APIRoot, path)
}

func printDump(d []byte, err error) {
	if err != nil {
		fmt.Printf("printDump err: %s\n", err)
		return
	}
	fmt.Printf("DUMP:\n%s\n", string(d))
}

func (c *Client) readResponse(resp *http.Response) (int, []byte, *map[string]string, error) {
	if resp == nil {
		return -1, nil, nil, fmt.Errorf("Resp is nil")
	}
	defer resp.Body.Close()
	if c.Debug {
		printDump(httputil.DumpResponse(resp, false))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, nil, err
	}
	respHeaders := &map[string]string{}
	for k, v := range resp.Header {
		(*respHeaders)[strings.ToLower(k)] = strings.Join(v, ",")
	}
	return resp.StatusCode, body, respHeaders, nil
}

func (c *Client) setReqHeaders(req *http.Request, headers *map[string]string) {
	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}
}

func (c *Client) get(path string, headers *map[string]string) (int, []byte, *map[string]string, error) {
	cli := c.httpClient()
	req, err := http.NewRequest("GET", c.composeURL(path), nil)
	if err != nil {
		return -1, nil, nil, err
	}
	req.Header.Add(headerAuthorization, c.constructAuthHeaderPAC())
	c.setReqHeaders(req, headers)
	if c.Debug {
		printDump(httputil.DumpRequest(req, false))
	}
	resp, err := cli.Do(req)
	if err != nil {
		return -1, nil, nil, err
	}
	return c.readResponse(resp)
}

func (c *Client) post(path string, body io.ReadCloser, headers *map[string]string) (int, []byte, *map[string]string, error) {
	if body != nil {
		defer body.Close()
	}
	cli := c.httpClient()
	req, err := http.NewRequest("POST", c.composeURL(path), body)
	if err != nil {
		return -1, nil, nil, err
	}
	req.Header.Add(headerAuthorization, c.constructAuthHeaderPAC())
	c.setReqHeaders(req, headers)
	if c.Debug {
		printDump(httputil.DumpRequest(req, false))
	}
	resp, err := cli.Do(req)
	if err != nil {
		return -1, nil, nil, err
	}
	return c.readResponse(resp)
}

func (c *Client) expect(code int, respBody []byte, expectedCode int, out interface{}, bin bool, binOut *[]byte, headers *map[string]string, outHeaders *map[string]string) *ScriveError {
	if code != expectedCode {
		se, err := c.parseResponseError(respBody)
		if err != nil {
			return localError(err)
		}
		return se
	}
	*outHeaders = *headers
	if bin {
		*binOut = respBody
		return nil
	}
	if out == nil {
		return nil
	}
	if err := parseJson(respBody, out); err != nil {
		return localError(err)
	}
	return nil
}

func (c *Client) getExpect(path string, headers *map[string]string, expectedCode int, out interface{}, bin bool, binOut *[]byte, outHeaders *map[string]string) *ScriveError {
	code, respBody, headers, err := c.get(path, headers)
	if err != nil {
		return localError(err)
	}
	return c.expect(code, respBody, expectedCode, out, bin, binOut, headers, outHeaders)
}

func (c *Client) postExpect(path string, body io.ReadCloser, headers *map[string]string, expectedCode int, out interface{}, bin bool, binOut *[]byte, outHeaders *map[string]string) *ScriveError {
	code, respBody, headers, err := c.post(path, body, headers)
	if err != nil {
		return localError(err)
	}
	return c.expect(code, respBody, expectedCode, out, bin, binOut, headers, outHeaders)
}

func (c *Client) doExpect(req *request) *ScriveError {
	if req.expectCode == noExpect {
		return localError(fmt.Errorf("Missing expect code"))
	}
	queryParams := req.getQuery()
	if queryParams != "" {
		queryParams = "?" + queryParams
	}
	switch req.method {
	case methodGET:
		return c.getExpect(
			req.path+queryParams,
			req.headers,
			req.expectCode,
			req.out,
			req.binaryResponse,
			&req.respBody,
			req.respHeaders,
		)
	case methodPOST:
		if err := req.finalize(); err != nil {
			return err
		}
		return c.postExpect(
			req.path+queryParams,
			ioutil.NopCloser(
				bytes.NewReader(req.reqBody),
			),
			req.headers,
			req.expectCode,
			req.out,
			req.binaryResponse,
			&req.respBody,
			req.respHeaders,
		)
	}
	return nil
}

func (c *Client) w(req *request, cb func(req *request)) (*request, *ScriveError) {
	if cb != nil {
		cb(req)
		if se := req.anyErrors(); se != nil {
			return nil, se
		}
	}
	return req, c.doExpect(req)
}

func (c *Client) we(req *request, path string, cb func(req *request), expectCode int, out interface{}) (*request, *ScriveError) {
	return c.w(req.Path(path).Expect(expectCode, out), cb)
}

func (c *Client) wb(req *request, path string, cb func(req *request), expectCode int) (*request, *ScriveError) {
	return c.w(req.Path(path).ExpectBinary(expectCode), cb)
}

func (c *Client) pwe(path string, cb func(req *request), expectCode int, out interface{}) (*request, *ScriveError) {
	return c.we(c.newPostRequest(), path, cb, expectCode, out)
}

func (c *Client) pw(path string, cb func(req *request), out interface{}) (*request, *ScriveError) {
	return c.pwe(path, cb, http.StatusOK, out)
}

func (c *Client) pwb(path string, cb func(req *request)) (*request, *ScriveError) {
	return c.wb(c.newPostRequest(), path, cb, http.StatusOK)
}

func (c *Client) pwbe(path string, cb func(req *request), expectCode int) (*request, *ScriveError) {
	return c.wb(c.newPostRequest(), path, cb, expectCode)
}

func (c *Client) gwe(path string, cb func(req *request), expectCode int, out interface{}) (*request, *ScriveError) {
	return c.we(c.newGetRequest(), path, cb, expectCode, out)
}

func (c *Client) gw(path string, cb func(req *request), out interface{}) (*request, *ScriveError) {
	return c.gwe(path, cb, http.StatusOK, out)
}

func (c *Client) gwb(path string, cb func(req *request)) (*request, *ScriveError) {
	return c.wb(c.newGetRequest(), path, cb, http.StatusOK)
}

func (c *Client) gwbe(path string, cb func(req *request), expectCode int) (*request, *ScriveError) {
	return c.wb(c.newGetRequest(), path, cb, expectCode)
}
