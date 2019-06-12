package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient        *http.Client
	Apikey            Key
	requestsPerSecond int
	testUrl           string
}

type Key struct {
	DistanceMatrixKey  string
	GoogleClientID     string
	GoogleClientSecret string
	OpenWeatherKey     string
	TankerkoenigKey    string
	HvvKey             string
}

// @Todo implement these two
var defaultRequestLimit = 50
var defaultRequestsPerSecond = 50

//@Todo add optional aruguments like below
// func NewClient(optional... Optional) (*Client, error)
func NewClient() *Client {
	c := &Client{requestsPerSecond: defaultRequestsPerSecond}
	return c
}

func NewClientWithTestUrl(testUrl string) *Client {
	c := &Client{}
	c.testUrl = testUrl
	return c
}

// Make a http GET Request to the Server
func (c *Client) get(config *apiConfig, apiReq apiRequest, resp interface{}) error {
	host := config.host
	//Only for testing purposes
	if c.testUrl != "" {
		host = c.testUrl
	}

	req, err := http.NewRequest(http.MethodGet, host+config.path, nil)
	if err != nil {
		return errors.New("failed to make a request to" + host + config.path)
	}
	if err := c.do(apiReq, req, resp); err != nil {
		return err
	}

	return nil
}

// Make a http POST Request to the Server
func (c *Client) post(config *apiConfig, apiReq apiRequest, resp interface{}, reqBody []byte, header map[string]string) error {
	host := config.host
	//Only for testing purposes
	if c.testUrl != "" {
		host = c.testUrl
	}

	req, err := http.NewRequest(http.MethodPost, host+config.path, bytes.NewBuffer(reqBody))
	if err != nil {
		return errors.New("failed to make a request to" + host + config.path)
	}

	//Set Headers
	for i, v := range header {
		req.Header.Set(i, v)
	}

	if err := c.do(apiReq, req, resp); err != nil {
		return err
	}

	return nil
}

func (c *Client) do(apiReq apiRequest, req *http.Request, resp interface{}) error {
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
		return fmt.Errorf("client failed to do request %v", err)
	}
	defer httpResp.Body.Close()

	return json.NewDecoder(httpResp.Body).Decode(&resp)

}

// Generates Urls for given Request
func (c *Client) generateUrls(q url.Values) string {
	return q.Encode()
}

// defines an interface for all API Requests
type apiRequest interface {
	params() url.Values
}
