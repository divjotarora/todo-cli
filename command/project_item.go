package command

import (
	"github.com/divjotarora/todo-cli/todoist"
	"github.com/elliotchance/orderedmap"
)

type projectListItem struct {
	*todoist.Project
	children *orderedmap.OrderedMap
}

var _ ListItem = (*projectListItem)(nil)

func newProjectListItem(project *todoist.Project) ListItem {
	return &projectListItem{
		Project:  project,
		children: orderedmap.NewOrderedMap(),
	}
}

func (p *projectListItem) Close() error {
	return nil
}

func (p *projectListItem) Children() *orderedmap.OrderedMap {
	return p.children
}

func (p *projectListItem) Parent() ListItem {
	return nil
}
