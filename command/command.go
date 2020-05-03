package command

import "github.com/divjotarora/todo-cli/api"

// Command repesents a command that can be executed.
type Command interface {
	Name() string
	Help() string
	Execute(ctx *Context, client *api.Client, args ...string) error
}
