package command

import (
	"github.com/divjotarora/todo-cli/todoist"
	"github.com/elliotchance/orderedmap"
)

type taskListItem struct {
	*todoist.Task

	parent   ListItem
	children *orderedmap.OrderedMap
}

var _ ListItem = (*taskListItem)(nil)

func newTaskListItem(task *todoist.Task, parent ListItem) ListItem {
	return &taskListItem{
		Task:     task,
		parent:   parent,
		children: orderedmap.NewOrderedMap(),
	}
}

func (t *taskListItem) Children() *orderedmap.OrderedMap {
	return t.children
}

func (t *taskListItem) Parent() ListItem {
	return t.parent
}
