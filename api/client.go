package api

import (
	"errors"
	"os"
)

// Client is an abstraction around the todoist API that can be used to make HTTP calls.
type Client struct {
	token string
}

// NewClient creates an API client.
func NewClient() (*Client, error) {
	token := os.Getenv("TODOIST_TOKEN")
	if token == "" {
		return nil, errors.New("TODOIST_TOKEN environment variable must be set")
	}

	c := &Client{
		token: token,
	}
	return c, nil
}
