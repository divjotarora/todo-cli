package app

import (
	"fmt"

	"github.com/desertbit/grumble"
	"github.com/divjotarora/todo-cli/command"
	"github.com/divjotarora/todo-cli/todoist"
)

// App is the main application that facilitates command line interactions.
type App struct {
	*grumble.App

	apiClient  *todoist.Client
	parser     *command.Parser
	commandCtx *command.Context
}

// NewApp creates a new command line application.
func NewApp() (*App, error) {
	apiClient, err := todoist.NewClient()
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
		cmd := cmd

		newApp.AddCommand(&grumble.Command{
			Name:      cmd.Name(),
			Help:      cmd.Help(),
			Usage:     cmd.Usage(),
			AllowArgs: true,
			Flags: func(grumbleFlags *grumble.Flags) {
				for _, flag := range cmd.Flags() {
					grumbleFlags.String(flag.ShortName, flag.LongName, flag.DefaultValue, flag.Description)
				}
			},
			Run: func(ctx *grumble.Context) error {
				var flagsMap map[string]string
				if len(ctx.Flags) > 0 {
					flagsMap = make(map[string]string)
					for key, value := range ctx.Flags {
						if stringValue, ok := value.Value.(string); ok {
							flagsMap[key] = stringValue
						}
					}
				}
				args := command.Arguments{
					PositionalArgs: ctx.Args,
					Flags:          flagsMap,
				}

				if err = cmd.Execute(newApp.commandCtx, newApp.apiClient, args); err != nil {
					return err
				}
				Display(newApp.commandCtx)
				return nil
			},
		})
	}

	newApp.OnInit(func(*grumble.App, grumble.FlagMap) error {
		projects, err := newApp.apiClient.Projects()
		if err != nil {
			return fmt.Errorf("error getting initial project list: %v", err)
		}

		newApp.commandCtx.SetProjectListing(projects)
		Display(newApp.commandCtx)
		return nil
	})

	return newApp, nil
}

// Start starts the application.
func (a *App) Start() {
	grumble.Main(a.App)
}
