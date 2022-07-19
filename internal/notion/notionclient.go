package notion

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	scheme        = "https"
	host          = "api.notion.com"
	apiVersion    = "v1"
	notionVersion = "2021-08-16"
)

type Client struct {
	httpClient    *http.Client
	urlScheme     string
	urlHost       string
	apiVersion    string
	notionVersion string
	Token         string
}

func NewClient(token string) (*Client, error) {

	if token == "" {
		err := errors.New("empty notion token")
		return nil, err
	}

	c := &Client{
		httpClient:    http.DefaultClient,
		urlScheme:     scheme,
		urlHost:       host,
		apiVersion:    apiVersion,
		notionVersion: notionVersion,
		Token:         token,
	}

	return c, nil
}

func (c *Client) request(id, method, object, requestBody string) (*http.Response, error) {

	var buf io.ReadWriter
	if requestBody != "" {
		buf = bytes.NewBuffer([]byte(requestBody))
	}

	requestURL, err := c.buildRequestURL(id, object)
	if err != nil {
		return nil, err
	}

	//requestURL := fmt.Sprintf("https://api.notion.com/%s/databases/%s/query", c.apiVersion, id)

	req, err := http.NewRequest(method, requestURL.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Add("Notion-Version", c.notionVersion)
	req.Header.Add("Content-Type", "application/json")

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("notion client response status %v", response.StatusCode))
	}

	return response, nil
}

func (c *Client) buildRequestURL(id, object string) (*url.URL, error) {
	var err error

	baseURL := &url.URL{
		Scheme: c.urlScheme,
		Host:   c.urlHost,
	}

	switch object {
	case "databases":
		baseURL, err = baseURL.Parse(fmt.Sprintf("%s/%s/%s/query", c.apiVersion, object, id))
		if err != nil {
			return nil, err
		}
	case "pages":
		baseURL, err = baseURL.Parse(fmt.Sprintf("%s/%s/%s", c.apiVersion, object, id))
		if err != nil {
			return nil, err
		}
	default:
		err := errors.New("missing object string in URL")
		return nil, err
	}

	return baseURL, nil
}
