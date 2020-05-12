package command

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/divjotarora/todo-cli/todoist"
	"github.com/elliotchance/orderedmap"
)

// ListItem TODO
type ListItem interface {
	Name() string
	ID() int64
	Close() error
	Children() *orderedmap.OrderedMap
	Parent() ListItem
}

// List TODO
type List struct {
	allItems *orderedmap.OrderedMap
}

func newList() *List {
	return &List{
		allItems: orderedmap.NewOrderedMap(),
	}
}

func newListFromTasks(tasks []*todoist.Task) *List {
	list := newList()
	for _, task := range tasks {
		taskItem := newTaskListItem(task, nil)
		list.allItems.Set(task.ID(), taskItem)

		for _, subtask := range task.Subtasks() {
			childItem := newTaskListItem(subtask, taskItem)
			taskItem.Children().Set(subtask.ID(), childItem)
		}
	}

	return list
}

func newListFromProjects(projects []*todoist.Project) *List {
	list := newList()
	for _, project := range projects {
		list.allItems.Set(project.ID(), newProjectListItem(project))
	}

	return list
}

// Get TODO
func (l *List) Get(index string) (ListItem, error) {
	indexes := strings.Split(index, ".")

	currentItems := l.allItems
	var selectedItem ListItem
	for _, indexStr := range indexes {
		convertedIdx, err := strconv.Atoi(indexStr)
		if err != nil || convertedIdx < 0 || convertedIdx >= currentItems.Len() {
			return nil, fmt.Errorf("invalid index %s", indexStr)
		}

		_, val, ok := currentItems.GetIndex(convertedIdx)
		if !ok {
			return nil, fmt.Errorf("invalid index %s", indexStr)
		}

		selectedItem = val.(ListItem)
		currentItems = selectedItem.Children()
	}
	return selectedItem, nil
}

// Delete TODO
func (l *List) Delete(item ListItem) {
	parent := item.Parent()

	parentMap := l.allItems
	if parent != nil {
		parentMap = parent.Children()
	}
	parentMap.Delete(item.ID())
}

// Items TODO
func (l *List) Items() *orderedmap.OrderedMap {
	return l.allItems
}
