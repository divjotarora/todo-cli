package app

import (
	"fmt"

	"github.com/desertbit/grumble"
	"github.com/divjotarora/todo-cli/api"
	"github.com/divjotarora/todo-cli/command"
)

// App is the main application that facilitates command line interactions.
type App struct {
	*grumble.App

	apiClient  *api.Client
	parser     *command.Parser
	commandCtx *command.Context
}

// NewApp creates a new command line application.
func NewApp() (*App, error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return nil, fmt.Errorf("error creating API client: %v", err)
	}

	newApp := &App{
		App: grumble.New(&grumble.Config{
			Name:        "todoist",
			Description: "Todoist CLI",
		}),

		apiClient:  apiClient,
		parser:     command.NewParser(),
		commandCtx: command.NewContext(),
	}

	for _, cmd := range newApp.parser.Commands() {
		newApp.AddCommand(&grumble.Command{
			Name: cmd.Name(),
			Help: cmd.Help(),
			Run: func(ctx *grumble.Context) error {
				return cmd.Execute(newApp.commandCtx, newApp.apiClient, ctx.Args...)
			},
		})
	}

	return newApp, nil
}

// Start starts the application.
func (a *App) Start() {
	grumble.Main(a.App)
}
