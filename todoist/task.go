package todoist

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

// Subtasks returns the sub-tasks for the task.
func (t *Task) Subtasks() []*Task {
	return t.subtasks
}
