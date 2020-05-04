package todoist

// Task represents a task in a project.
type Task struct {
	Name     string
	ID       int64
	Subtasks []Task

	client *Client
}
