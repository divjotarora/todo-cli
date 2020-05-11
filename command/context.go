package command

import (
	"github.com/divjotarora/todo-cli/todoist"
)

// Context contains information about the application state that can be used when executing a command.
type Context struct {
	selectedProject *todoist.Project
	currentListing  *List
}

// NewContext creates a new context.
func NewContext() *Context {
	return &Context{}
}

// SetProject sets the value for the current working project.
func (c *Context) SetProject(project *todoist.Project) {
	c.selectedProject = project
}

// SetProjectListing sets the current listing to a slice of projects.
func (c *Context) SetProjectListing(projects []*todoist.Project) {
	c.currentListing = newListFromProjects(projects)
}

// SetTaskListing sets the current listing to a slice of tasks.
func (c *Context) SetTaskListing(tasks []*todoist.Task) {
	c.currentListing = newListFromTasks(tasks)
}

// CurrentListing returns the current list of items associated with the Context.
func (c *Context) CurrentListing() *List {
	return c.currentListing
}
