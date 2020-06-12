package todoist

import (
	"encoding/json"
	"fmt"
)

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

// ID returns the project's ID.
func (p *Project) ID() int64 {
	return p.id
}

type unmarshalTask struct {
	ID      int64
	Parent  *int64
	Content string
}

func (p *Project) newTask(temp *unmarshalTask) *Task {
	return &Task{
		name:     temp.Content,
		id:       temp.ID,
		parentID: temp.Parent,
		client:   p.client,
	}
}

// Tasks returns the list of tasks in this project.
func (p *Project) Tasks() ([]*Task, error) {
	endpoint := fmt.Sprintf("tasks?project_id=%d", p.id)
	res, err := p.client.httpGet(endpoint)
	if err != nil {
		return nil, newError(err)
	}

	var unmarshalled []*unmarshalTask
	if err = json.Unmarshal(res, &unmarshalled); err != nil {
		return nil, newError(err)
	}

	// Do one iteration to convert the unmarshalled slice into a slice of Task instances and find all of the top-level
	// tasks.
	tasksMap := make(map[int64]*Task)
	var tasks []*Task
	for _, temp := range unmarshalled {
		newTask := p.newTask(temp)
		tasksMap[newTask.id] = newTask

		if temp.Parent == nil {
			tasks = append(tasks, newTask)
		}
	}

	// Do another iteration to add subtasks to their parent's subtasks slice. This has to be done separately from the
	// iteration above because the parent task might appear after the subtask in the API response.
	for _, task := range tasksMap {
		if task.parentID == nil {
			continue
		}

		parent := tasksMap[*task.parentID]
		parent.subtasks = append(parent.subtasks, task)
	}

	return tasks, nil
}

// CreateTask creates a new task in this project.
func (p *Project) CreateTask(content string) (*Task, error) {
	return p.createTask(content, nil)
}

// CreateSubTask creates a new subtask under the provided parent ID.
func (p *Project) CreateSubTask(content string, parentTaskID int64) (*Task, error) {
	return p.createTask(content, &parentTaskID)
}

func (p *Project) createTask(content string, parent *int64) (*Task, error) {
	postBody := map[string]interface{}{
		"content":    content,
		"project_id": p.id,
	}
	if parent != nil {
		postBody["parent"] = *parent
	}

	res, err := p.client.httpPost("tasks", postBody)
	if err != nil {
		return nil, newError(err)
	}

	var unmarshalled *unmarshalTask
	if err = json.Unmarshal(res, &unmarshalled); err != nil {
		return nil, newError(err)
	}
	return p.newTask(unmarshalled), nil
}
