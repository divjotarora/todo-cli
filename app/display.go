package app

import (
	"fmt"

	"github.com/divjotarora/todo-cli/command"
)

// Display prints a list of Nameable instances to the console.
func Display(ctx *command.Context) {
	displayHelper(ctx.CurrentListing(), "")
}

func displayHelper(items []command.ListItem, prefix string) {
	for idx, item := range items {
		fmt.Printf("%v%v. %v\n", prefix, idx, item.Name())

		nested, ok := item.(command.NestedListItem)
		if !ok {
			continue
		}

		if children := nested.Children(); len(children) > 0 {
			displayHelper(children, prefix+"  ")
		}
	}
}
