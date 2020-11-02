package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
    _"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/go-cleanhttp"
)

// Client is a Grafana API client.
type Client struct {
	config  Config
	baseURL url.URL
	client  *http.Client
}

// Config contains client configuration.
type Config struct {
	// APIKey is an optional API key.
	APIKey string
	// BasicAuth is optional basic auth credentials.
	BasicAuth *url.Userinfo
	// Client provides an optional HTTP client, otherwise a default will be used.
	Client *http.Client
}

// New creates a new Grafana client.
func New(baseURL string, cfg Config) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if cfg.BasicAuth != nil {
		u.User = cfg.BasicAuth
	}

	cli := cfg.Client
	if cli == nil {
		cli = cleanhttp.DefaultClient()
	}

	return &Client{
		config:  cfg,
		baseURL: *u,
		client:  cli,
	}, nil
}

func (c *Client) request(method, requestPath string, query url.Values, body io.Reader, responseStruct interface{}) error {
    headers := make(map[string]string, 0)

	r, err := c.newRequest(method, requestPath, query, body, headers)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyContents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if os.Getenv("GF_LOG") != "" {
		log.Printf("response status %d with body %v", resp.StatusCode, string(bodyContents))
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("status: %d, body: %v", resp.StatusCode, string(bodyContents))
	}

	if responseStruct == nil {
		return nil
	}

	err = json.Unmarshal(bodyContents, responseStruct)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) requestWithHeaders(method, requestPath string, query url.Values, body io.Reader, responseStruct interface{}, headers map[string]string) error {
	r, err := c.newRequest(method, requestPath, query, body, headers)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyContents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if os.Getenv("GF_LOG") != "" {
		log.Printf("response status %d with body %v", resp.StatusCode, string(bodyContents))
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("status: %d, body: %v", resp.StatusCode, string(bodyContents))
	}

	if responseStruct == nil {
		return nil
	}

	err = json.Unmarshal(bodyContents, responseStruct)
	if err != nil {
		return err
	}

	return nil
}



func (c *Client) newRequest(method, requestPath string, query url.Values, body io.Reader, headers map[string]string) (*http.Request, error) {
	url := c.baseURL
	url.Path = path.Join(url.Path, requestPath)
	url.RawQuery = query.Encode()
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return req, err
	}

	if c.config.APIKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.APIKey))
	}

	if os.Getenv("GF_LOG") != "" {
		if body == nil {
			log.Printf("request (%s) to %s with no body data", method, url.String())
		} else {
			log.Printf("request (%s) to %s with body data: %s", method, url.String(), body.(*bytes.Buffer).String())
		}
	}

    // add extra headers
    for h, v := range headers {
        req.Header.Add(h, v)
    }

	req.Header.Add("Content-Type", "application/json")

	return req, err
}
