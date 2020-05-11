package command

import (
	"github.com/cevaris/ordered_map"
	"github.com/divjotarora/todo-cli/todoist"
)

type projectListItem struct {
	*todoist.Project
	children *ordered_map.OrderedMap
}

var _ ListItem = (*projectListItem)(nil)

func newProjectListItem(project *todoist.Project) *projectListItem {
	return &projectListItem{
		Project:  project,
		children: ordered_map.NewOrderedMap(),
	}
}

func (p *projectListItem) Close() error {
	return nil
}

func (p *projectListItem) Children() *ordered_map.OrderedMap {
	return p.children
}

func (p *projectListItem) Parent() ListItem {
	return nil
}
