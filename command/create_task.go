package command

import "github.com/divjotarora/todo-cli/todoist"

// CreateTask is the command to create a new task in a project.
var CreateTask Command = &createTask{}

type createTask struct{}

func (c *createTask) Name() string {
	return "create"
}

func (c *createTask) Help() string {
	return "create a new task"
}

func (c *createTask) Usage() string {
	return "create <content> -s <subtask parent>"
}

func (c *createTask) Flags() []Flag {
	subtaskFlag := Flag{
		ShortName:   "s",
		LongName:    "subtask",
		Description: "the index of the parent issue",
	}

	return []Flag{subtaskFlag}
}

func (c *createTask) Execute(ctx *Context, client *todoist.Client, args Arguments) error {
	return nil
}
