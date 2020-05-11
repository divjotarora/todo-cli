package command

import (
	"github.com/cevaris/ordered_map"
	"github.com/divjotarora/todo-cli/todoist"
)

type taskListItem struct {
	*todoist.Task

	parent   *taskListItem
	children *ordered_map.OrderedMap
}

var _ ListItem = (*taskListItem)(nil)

func newTaskListItem(task *todoist.Task, parent *taskListItem) *taskListItem {
	return &taskListItem{
		Task:     task,
		parent:   parent,
		children: ordered_map.NewOrderedMap(),
	}
}

func (t *taskListItem) Children() *ordered_map.OrderedMap {
	return t.children
}

func (t *taskListItem) Parent() ListItem {
	return t.parent
}
