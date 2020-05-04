package command

import "github.com/divjotarora/todo-cli/todoist"

// Flag represents a command line argument flag.
type Flag struct {
	ShortName    string
	LongName     string
	DefaultValue string
	Description  string
}

// Arguments represents the inputs for a command.
type Arguments struct {
	PositionalArgs []string
	Flags          map[string]string
}

// Command repesents a command that can be executed.
type Command interface {
	Name() string
	Help() string
	Usage() string
	Flags() []Flag
	Execute(*Context, *todoist.Client, Arguments) error
}
