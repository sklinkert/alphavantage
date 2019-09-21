package alphavantage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const baseURL = "https://www.alphavantage.co"

// Client represents a new alphavantage client
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// New creates new Client instance
func New(apiKey string) *Client {
	const httpTimeout = time.Second * 30

	httpClient := &http.Client{
		Timeout: httpTimeout,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 5,
		},
	}

	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (c *Client) makeHTTPRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("building http request failed: %w", err)
	}
	req.Header.Set("User-Agent", "Go client: github.com/sklinkert/alphavantage")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: expected %d, got %d",
			http.StatusOK, resp.StatusCode)
	}

	return body, nil
}
