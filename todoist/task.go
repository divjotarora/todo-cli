package todoist

import "fmt"

// Task represents a task in a project.
type Task struct {
	name     string
	id       int64
	subtasks []*Task
	client   *Client
}

// Name returns the name of the task.
func (t *Task) Name() string {
	return t.name
}

// ID returns the task ID.
func (t *Task) ID() int64 {
	return t.id
}

// Subtasks returns the sub-tasks for the task.
func (t *Task) Subtasks() []*Task {
	return t.subtasks
}

// Close closes a task.
func (t *Task) Close() error {
	endpoint := fmt.Sprintf("tasks/%d/close", t.id)
	_, err := t.client.httpPost(endpoint, nil)
	if err != nil {
		return newError(err)
	}

	return nil
}
