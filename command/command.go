package command

import "github.com/divjotarora/todo-cli/todoist"

// Flag represents a command line argument flag.
type Flag struct {
	ShortName    string
	LongName     string
	DefaultValue string
	Description  string
}

// Flags represents an interface to retrieve command line flags. Each method corresponds a Go type. The implementation
// should return the flag value corresponding to the provided flag name as that type.
type Flags interface {
	String(flagName string) string
}

// Arguments represents the inputs for a command.
type Arguments struct {
	PositionalArgs []string
	Flags          Flags
}

// Command repesents a command that can be executed.
type Command interface {
	Name() string
	Help() string
	Usage() string
	Flags() []Flag
	Execute(*Context, *todoist.Client, Arguments) error
}
