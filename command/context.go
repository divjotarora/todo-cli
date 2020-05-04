package command

import (
	"github.com/divjotarora/todo-cli/todoist"
)

// ListItem represents an item in a listing.
type ListItem interface {
	Name() string
}

// NestedListItem represents an item in a listing that can have one or more sub-items.
type NestedListItem interface {
	ListItem
	Children() []ListItem
}

// Context contains information about the application state that can be used when executing a command.
type Context struct {
	selectedProject *todoist.Project
	currentListing  []ListItem
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
	items := make([]ListItem, 0, len(projects))
	for _, p := range projects {
		items = append(items, p)
	}

	c.currentListing = items
}

// SetTaskListing sets the current listing to a slice of tasks.
func (c *Context) SetTaskListing(tasks []*todoist.Task) {
	c.currentListing = newTaskListItems(tasks)
}

// CurrentListing returns the current list of items associated with the Context.
func (c *Context) CurrentListing() []ListItem {
	return c.currentListing
}
