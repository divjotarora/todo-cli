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
	nameable := make([]ListItem, 0, len(projects))
	for _, p := range projects {
		nameable = append(nameable, p)
	}

	c.currentListing = nameable
}

// CurrentListing returns the current list of items associated with the Context.
func (c *Context) CurrentListing() []ListItem {
	return c.currentListing
}
