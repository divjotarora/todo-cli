package command

import "github.com/divjotarora/todo-cli/todoist"

type taskListItem struct {
	*todoist.Task
}

var _ ListItem = (*taskListItem)(nil)
var _ NestedListItem = (*taskListItem)(nil)

func (t *taskListItem) Children() []ListItem {
	subtasks := t.Subtasks()
	if len(subtasks) == 0 {
		return nil
	}

	return newTaskListItems(subtasks)
}

func newTaskListItems(tasks []*todoist.Task) []ListItem {
	items := make([]ListItem, 0, len(tasks))
	for _, task := range tasks {
		items = append(items, &taskListItem{task})
	}
	return items
}
