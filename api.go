package scrive

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
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

func (c *Client) readResponse(resp *http.Response) (int, []byte, error) {
	if resp == nil {
		return -1, nil, fmt.Errorf("Resp is nil")
	}
	defer resp.Body.Close()
	printDump(httputil.DumpResponse(resp, false))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, err
	}
	return resp.StatusCode, body, nil
}

func (c *Client) get(path string) (int, []byte, error) {
	cli := c.httpClient()
	req, err := http.NewRequest("GET", c.composeURL(path), nil)
	if err != nil {
		return -1, nil, err
	}
	req.Header.Add(headerAuthorization, c.constructAuthHeaderPAC())
	printDump(httputil.DumpRequest(req, false))
	resp, err := cli.Do(req)
	if err != nil {
		return -1, nil, err
	}
	return c.readResponse(resp)
}

func (c *Client) post(path string, body io.ReadCloser, headers *map[string]string) (int, []byte, error) {
	if body != nil {
		defer body.Close()
	}
	cli := c.httpClient()
	req, err := http.NewRequest("POST", c.composeURL(path), body)
	if err != nil {
		return -1, nil, err
	}
	req.Header.Add(headerAuthorization, c.constructAuthHeaderPAC())
	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}
	printDump(httputil.DumpRequest(req, false))
	resp, err := cli.Do(req)
	if err != nil {
		return -1, nil, err
	}
	return c.readResponse(resp)
}

func (c *Client) postExpect(path string, body io.ReadCloser, headers *map[string]string, expectedCode int, out interface{}) *ScriveError {
	code, respBody, err := c.post(path, body, headers)
	if err != nil {
		return localError(err)
	}
	if code != expectedCode {
		se, err := c.parseResponseError(respBody)
		if err != nil {
			return localError(err)
		}
		return se
	}
	if err := parseJson(respBody, out); err != nil {
		return localError(err)
	}
	return nil
}
