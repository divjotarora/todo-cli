package command

import "github.com/divjotarora/todo-cli/todoist"

// CreateTask is the command to create a new task in a project.
var CreateTask Command = &createTask{}

type createTask struct{}

func (c *createTask) Name() string {
	return "create"
}

func (c *createTask) Help() string {
	return "create a new task"
}

func (c *createTask) Usage() string {
	return "create <content> -s <subtask parent>"
}

func (c *createTask) Flags() []Flag {
	subtaskFlag := Flag{
		ShortName:   "s",
		LongName:    "subtask",
		Description: "the index of the parent issue",
	}

	return []Flag{subtaskFlag}
}

func (c *createTask) Execute(ctx *Context, client *todoist.Client, args Arguments) error {
	if len(args.PositionalArgs) != 1 {
		return invalidNumberOfArguments(c, 1, len(args.PositionalArgs))
	}

	project := ctx.selectedProject
	content := args.PositionalArgs[0]
	var parentTask ListItem
	var created *todoist.Task
	var err error

	if subtaskIndex := args.Flags.String("subtask"); subtaskIndex != "" {
		parentTask, err = ctx.currentListing.Get(subtaskIndex)
		if err != nil {
			return err
		}

		created, err = project.CreateSubTask(content, parentTask.ID())
	} else {
		created, err = project.CreateTask(content)
	}
	if err != nil {
		return err
	}

	ctx.AddTask(created, parentTask)
	return nil
}
