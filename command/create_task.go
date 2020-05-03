package command

import "github.com/divjotarora/todo-cli/api"

// CreateTask is the command to create a new task in a project.
var CreateTask Command = &createTask{}

type createTask struct{}

func (c *createTask) Name() string {
	return "createTask"
}

func (c *createTask) Help() string {
	return "create a new task"
}

func (c *createTask) Execute(ctx *Context, apiClient *api.Client, args ...string) error {
	return nil
}
