package command

import (
	"github.com/divjotarora/todo-cli/todoist"
)

// GetTasks is the command to get the tasks for the current working project.
var GetTasks Command = &getTasks{}

type getTasks struct{}

func (*getTasks) Name() string {
	return "tasks"
}

func (*getTasks) Help() string {
	return "get a list of tasks for the current working project"
}

func (*getTasks) Usage() string {
	return "tasks"
}

func (*getTasks) Flags() []Flag {
	return nil
}

func (s *getTasks) Execute(ctx *Context, client *todoist.Client, args Arguments) error {
	if len(args.PositionalArgs) != 0 {
		return invalidNumberOfArguments(s, 0, len(args.PositionalArgs))
	}

	tasks, err := ctx.selectedProject.Tasks()
	if err != nil {
		return err
	}

	ctx.SetTaskListing(tasks)
	return nil
}
