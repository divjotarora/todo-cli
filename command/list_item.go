package command

import (
	"github.com/cevaris/ordered_map"
	"github.com/divjotarora/todo-cli/todoist"
)

// ListItem TODO
type ListItem interface {
	Name() string
	ID() int64
	Close() error
	Children() *ordered_map.OrderedMap
	Parent() ListItem
}

// List TODO
type List struct {
	allItems *ordered_map.OrderedMap
}

func newList() *List {
	return &List{
		allItems: ordered_map.NewOrderedMap(),
	}
}

func newListFromTasks(tasks []*todoist.Task) *List {
	list := newList()
	for _, task := range tasks {
		taskItem := newTaskListItem(task, nil)
		list.allItems.Set(task.ID(), taskItem)

		for _, subtask := range task.Subtasks() {
			childItem := newTaskListItem(subtask, taskItem)
			taskItem.children.Set(subtask.ID(), childItem)
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
func (l *List) Items() *ordered_map.OrderedMap {
	return l.allItems
}

// // Iterator TODO
// type Iterator struct {
// 	iter func() (*ordered_map.KVPair, bool)
// 	done bool
// }

// // Next TODO
// func (i *Iterator) Next() ListItem {
// 	if i.done {
// 		return nil
// 	}

// 	kv, ok := i.iter()
// 	if !ok {
// 		return nil
// 	}
// 	return kv.Value.(ListItem)
// }

// // Iterator TODO
// func (l *List) Iterator() *Iterator {
// 	return &Iterator{
// 		iter: l.allItems.IterFunc(),
// 	}
// }
