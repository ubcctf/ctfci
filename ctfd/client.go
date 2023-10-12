package ctfd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIClient struct {
	Token string
	URL   string
	http.Client
}

// NewAPIClient creates an api client with auth token and CTFd url without trailing slash.
func NewAPIClient(token string, url string) *APIClient {
	return &APIClient{
		Token: token,
		URL:   url,
	}
}

func (c *APIClient) request(method string, url string, reqBody io.Reader) (map[string]interface{}, error) {
	req, err := http.NewRequest(method, c.URL+url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %v", c.Token))
	req.Header.Add("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request failed %v", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	json.Unmarshal(body, &data)
	return data, nil
}

func (c *APIClient) get(url string) (map[string]interface{}, error) {
	return c.request("GET", url, nil)
}

func (c *APIClient) post(url string, body io.Reader) (map[string]interface{}, error) {
	return c.request("POST", url, body)
}

func (c *APIClient) patch(url string, body io.Reader) (map[string]interface{}, error) {
	return c.request("PATCH", url, body)
}

func (c *APIClient) delete(url string) (map[string]interface{}, error) {
	return c.request("DELETE", url, nil)
}
