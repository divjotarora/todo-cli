package command

import "github.com/divjotarora/todo-cli/todolist"

// Context contains information about the application state that can be used when executing a command.
type Context struct {
	currentProjects []todolist.Project
	currentTasks    []todolist.Task
}

// NewContext creates a new context.
func NewContext() *Context {
	return &Context{}
}
