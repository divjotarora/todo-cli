package todoist

import (
	"encoding/json"
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

type unmarshalProject struct {
	Name string
	ID   int64
}

func (c *Client) newProject(temp unmarshalProject) *Project {
	return &Project{
		name:   temp.Name,
		id:     temp.ID,
		client: c,
	}
}

// Projects gets a list of all projects.
func (c *Client) Projects() ([]*Project, error) {
	res, err := c.httpGet("projects")
	if err != nil {
		return nil, newError(err)
	}

	var unmarshalled []unmarshalProject
	if err = json.Unmarshal(res, &unmarshalled); err != nil {
		return nil, newError(err)
	}

	projects := make([]*Project, 0, len(unmarshalled))
	for _, temp := range unmarshalled {
		projects = append(projects, c.newProject(temp))
	}
	return projects, nil
}
