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

type unmarshalTask struct {
	ID      int64
	Parent  *int64
	Content string
}

func (p *Project) newTask(temp *unmarshalTask) *Task {
	return &Task{
		name:   temp.Content,
		id:     temp.ID,
		client: p.client,
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

	tasksMap := make(map[int64]*Task)
	var tasks []*Task
	for _, temp := range unmarshalled {
		newTask := p.newTask(temp)
		tasksMap[newTask.id] = newTask

		if temp.Parent == nil {
			tasks = append(tasks, newTask)
			continue
		}

		parent := tasksMap[*temp.Parent]
		parent.subtasks = append(parent.subtasks, newTask)
	}

	return tasks, nil
}
