package todoist

// Project represents a project in the todolist.
type Project struct {
	name   string
	id     int64
	client *Client
}

// Name returns the project's name.
func (p *Project) Name() string {
	return p.name
}
