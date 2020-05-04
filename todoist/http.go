package todoist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

const (
	apiURL = "https://api.todoist.com/rest/v1"
)

func (c *Client) httpRequest(method, endpoint string, body map[string]interface{}) ([]byte, error) {
	httpClient := &http.Client{}

	var bodyBuffer io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		bodyBuffer = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequest(string(method), fmt.Sprintf("%s/%s", apiURL, endpoint), bodyBuffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Request-Id", uuid.New().String())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	if len(body) != 0 {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 && res.StatusCode != 204 {
		return nil, fmt.Errorf("API error: %s", respBody)
	}
	return respBody, nil
}

func (c *Client) httpGet(endpoint string) ([]byte, error) {
	return c.httpRequest("GET", endpoint, nil)
}

func (c *Client) httpPost(endpoint string, body map[string]interface{}) ([]byte, error) {
	return c.httpRequest("POST", endpoint, body)
}

func (c *Client) httpDelete(endpoint string, body map[string]interface{}) ([]byte, error) {
	return c.httpRequest("DELETE", endpoint, body)
}
