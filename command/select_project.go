package command

import (
	"fmt"
	"strconv"

	"github.com/divjotarora/todo-cli/todoist"
)

// SelectProject is the command to change the current working project.
var SelectProject Command = &selectProject{}

type selectProject struct{}

func (*selectProject) Name() string {
	return "select"
}

func (*selectProject) Help() string {
	return "select a project to be the current working project"
}

func (*selectProject) Usage() string {
	return "select <project index>"
}

func (*selectProject) Flags() []Flag {
	return nil
}

func (s *selectProject) Execute(ctx *Context, client *todoist.Client, args Arguments) error {
	if len(args.PositionalArgs) != 1 {
		return invalidNumberOfArguments(s, 1, len(args.PositionalArgs))
	}

	projects, err := client.Projects()
	if err != nil {
		return err
	}

	index, err := strconv.Atoi(args.PositionalArgs[0])
	if err != nil {
		return newExecutionError(s, err)
	}
	if index >= len(projects) {
		wrapped := fmt.Errorf("index %d is longer than the list of projects", index)
		return newExecutionError(s, wrapped)
	}

	ctx.SetProject(projects[index])
	return GetTasks.Execute(ctx, client, Arguments{})
}
