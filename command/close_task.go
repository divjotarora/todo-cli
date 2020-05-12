package command

import (
	"fmt"

	"github.com/divjotarora/todo-cli/todoist"
)

// CloseTask is the command to close a task in a project.
var CloseTask Command = &closeTask{}

type closeTask struct{}

func (c *closeTask) Name() string {
	return "close"
}

func (c *closeTask) Help() string {
	return "close a task"
}

func (c *closeTask) Usage() string {
	return "close <index 1> <index 2>..."
}

func (c *closeTask) Flags() []Flag {
	return nil
}

func (c *closeTask) Execute(ctx *Context, client *todoist.Client, args Arguments) error {
	if len(args.PositionalArgs) == 0 {
		return invalidNumberOfArguments(c, 1, len(args.PositionalArgs))
	}

	items := make([]ListItem, 0, len(args.PositionalArgs))
	for _, arg := range args.PositionalArgs {
		item, err := ctx.currentListing.Get(arg)
		if err != nil {
			return err
		}

		if err = item.Close(); err != nil {
			return fmt.Errorf("error closing item %s: %v", item.Name(), err)
		}

		items = append(items, item)
	}
	for _, item := range items {
		ctx.currentListing.Delete(item)
	}
	return nil
}
