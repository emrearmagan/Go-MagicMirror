package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient        *http.Client
	name              string
	requestsPerSecond int

	TestUrl string
}

// @Todo implement these two
var defaultRequestLimit = 50
var defaultRequestsPerSecond = 50

//@Todo add optional aruguments like below
// func NewClient(optional... Optional) (*Client, error)
func NewClient() (*Client, error) {
	c := &Client{requestsPerSecond: defaultRequestsPerSecond}
	return c, nil
}

func NewClientWithTestUrl(TestUrl string) (*Client, error) {
	c := &Client{requestsPerSecond: defaultRequestsPerSecond}
	c.TestUrl = TestUrl
	return c, nil
}

// Make a http GET Request to the Server
func (c *Client) get(config *ApiConfig, apiReq apiRequest, resp interface{}) error {
	host := config.host
	//Only for testing purposes
	if c.TestUrl != "" {
		host = c.TestUrl
	}

	req, err := http.NewRequest(http.MethodGet, host+config.path, nil)
	if err != nil {
		return errors.New("failed to make a request to" + host + config.path)
	}
	if err := c.do(config, apiReq, req, resp); err != nil {
		return err
	}

	return nil
}

// Make a http POST Request to the Server
func (c *Client) post(config *ApiConfig, apiReq apiRequest, resp interface{}, reqBody []byte, header map[string]string) error {
	host := config.host
	//Only for testing purposes
	if c.TestUrl != "" {
		host = c.TestUrl
	}

	req, err := http.NewRequest(http.MethodPost, host+config.path, bytes.NewBuffer(reqBody))
	if err != nil {
		return errors.New("failed to make a request to" + host + config.path)
	}

	//Set Headers
	for i, v := range header {
		req.Header.Set(i, v)
	}

	if err := c.do(config, apiReq, req, resp); err != nil {
		return err
	}

	return nil
}

func (c *Client) do(config *ApiConfig, apiReq apiRequest, req *http.Request, resp interface{}) error {
	if apiReq != nil {
		urls := c.generateUrls(apiReq.params())
		req.URL.RawQuery = urls
	}

	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}

	httpResp, err := client.Do(req)
	if err != nil {
		return errors.New("client failed to do request")
	}
	defer httpResp.Body.Close()

	return json.NewDecoder(httpResp.Body).Decode(resp)
}

// Generates Urls for given Request
func (c *Client) generateUrls(q url.Values) string {
	return q.Encode()
}

// defines an interface for all API Requests
type apiRequest interface {
	params() url.Values
}
