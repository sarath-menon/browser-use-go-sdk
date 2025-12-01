package browseruse

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	defaultBaseURL = "https://api.browser-use.com/api/v2"
	defaultTimeout = 30 * time.Second
)

// Client is the main Browser Use API client
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	Tasks      *TasksService
}

// ClientOptions configures the Browser Use client
type ClientOptions struct {
	// APIKey for authentication. If empty, reads from BROWSER_USE_API_KEY env var
	APIKey string
	// BaseURL for the API. Defaults to https://api.browser-use.com/api/v2
	BaseURL string
	// HTTPClient to use for requests. If nil, uses default client
	HTTPClient *http.Client
	// Timeout for HTTP requests. Defaults to 30 seconds
	Timeout time.Duration
}

// NewClient creates a new Browser Use API client
func NewClient(opts *ClientOptions) (*Client, error) {
	if opts == nil {
		opts = &ClientOptions{}
	}

	apiKey := opts.APIKey
	if apiKey == "" {
		apiKey = os.Getenv("BROWSER_USE_API_KEY")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required: provide via options or BROWSER_USE_API_KEY env var")
	}

	baseURL := opts.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	httpClient := opts.HTTPClient
	if httpClient == nil {
		timeout := opts.Timeout
		if timeout == 0 {
			timeout = defaultTimeout
		}
		httpClient = &http.Client{
			Timeout: timeout,
		}
	}

	client := &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: httpClient,
	}

	client.Tasks = &TasksService{client: client}

	return client, nil
}

// doRequest executes an HTTP request and decodes the response
func (c *Client) doRequest(ctx context.Context, method, path string, body io.Reader, result interface{}) error {
	url := c.baseURL + path

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Browser-Use-API-Key", c.apiKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return parseError(resp.StatusCode, respBody)
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
